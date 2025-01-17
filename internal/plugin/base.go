package plugin

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/LK4D4/joincontext"
	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/cacher"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/cleanup"
	"github.com/hashicorp/vagrant-plugin-sdk/internal-shared/dynamic"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/funcspec"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/pluginargs"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
)

// Checks if value is implemented. If it is not (value is nil) a
// customized error is returned
func isImplemented(
	t interface{}, // value to check
	name string, // name of the value
) error {
	if t == nil {
		return status.Errorf(codes.Unimplemented, "plugin does not implement: "+name)
	}
	return nil
}

// Defines client interface for plugins supporting value seeding
type SeederClient interface {
	Seed(ctx context.Context, in *vagrant_plugin_sdk.Args_Seeds, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Seeds(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vagrant_plugin_sdk.Args_Seeds, error)
}

// Defines client interface for plugins supporting a name
type NamedPluginClient interface {
	SetPluginName(ctx context.Context, in *vagrant_plugin_sdk.PluginInfo_Name, opts ...grpc.CallOption) (*emptypb.Empty, error)
	PluginName(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*vagrant_plugin_sdk.PluginInfo_Name, error)
}

// BasePlugin contains the information which is common among
// all plugins. It should be embedded in every plugin type.
type BasePlugin struct {
	Cache   cacher.Cache      // Cache for mappers
	Cleanup cleanup.Cleanup   // Used to register cleanup tasks
	Mappers []*argmapper.Func // Mappers
	Logger  hclog.Logger      // Logger
	Wrapped bool              // Used to determine if wrapper
}

// Base client type
type BaseClient struct {
	*Base

	Ctx             context.Context // base context for the client
	Client          interface{}     // actual grpc client
	addr            net.Addr        // address to connect to this client
	parentComponent interface{}     // parent component (if available)
}

// Base server type
type BaseServer struct {
	*Base

	impl       interface{}                    // real implementation
	name       string                         // name of the plugin
	seedValues *vagrant_plugin_sdk.Args_Seeds // stored seed values
}

// Create a new shallow copy
func (b *BasePlugin) Clone() *BasePlugin {
	return &BasePlugin{
		Cache:   b.Cache,
		Cleanup: cleanup.New(),
		Mappers: mappers(b.Mappers),
		Logger:  b.Logger,
		Wrapped: b.Wrapped,
	}
}

// Create a new client
func (b *BasePlugin) NewClient(
	ctx context.Context, // context used by the client
	broker *plugin.GRPCBroker, // broker assigned to this client
	s interface{}, // the grpc client
) *BaseClient {
	return &BaseClient{
		Ctx:    ctx,
		Client: s,
		Base: &Base{
			Broker:  broker,
			Cache:   b.Cache,
			Cleanup: cleanup.New(),
			Logger:  b.Logger,
			Mappers: mappers(b.Mappers),
			Wrapped: b.Wrapped,
		},
	}
}

// Create a new server
func (b *BasePlugin) NewServer(
	broker *plugin.GRPCBroker, // broker assigned to this server
	impl interface{}, // real implementation of service
) *BaseServer {
	return &BaseServer{
		impl:       impl,
		seedValues: &vagrant_plugin_sdk.Args_Seeds{},
		name:       "",
		Base: &Base{
			Broker:  broker,
			Cache:   b.Cache,
			Cleanup: cleanup.New(),
			Logger:  b.Logger,
			Mappers: mappers(b.Mappers),
			Wrapped: b.Wrapped,
		},
	}
}

// Base contains shared logic for all plugin server/client implementations.
// This should be embedded in every plugin server/client implementation using
// the specialized server and client types.
type Base struct {
	Broker  *plugin.GRPCBroker
	Logger  hclog.Logger
	Mappers []*argmapper.Func
	Cleanup cleanup.Cleanup
	Cache   cacher.Cache
	Wrapped bool
}

// Create a new BasePlugin that is a shallow copy
// of the current Base and marks wrapped as true.
// This is used when wrapping a GRPC client.
func (b *Base) Wrap() *BasePlugin {
	return &BasePlugin{
		Cache:   b.Cache,
		Cleanup: cleanup.New(),
		Logger:  b.Logger,
		Mappers: mappers(b.Mappers),
		Wrapped: true,
	}
}

// If this plugin is a wrapper
func (b *Base) IsWrapped() bool {
	return b.Wrapped
}

// internal returns a new pluginargs.Internal that can be used with
// dynamic calls. The Internal structure is an internal-only argument
// that is used to perform cleanup.
func (b *Base) Internal() pluginargs.Internal {
	// if the cache isn't currently set, just create
	// a new cache instance and set it now
	if b.Cache == nil {
		b.Cache = cacher.New()
	}

	return pluginargs.New(
		b.Broker,
		b.Cache,
		b.Cleanup,
		b.Logger,
		b.Mappers,
	)
}

// Map a value to the expected type using registered mappers
// NOTE: The expected type must be a pointer, so an expected type
// of `*int` means an `int` is wanted. Expected type of `**int`
// means an `*int` is wanted, etc.
func (b *Base) Map(
	resultValue, // value to be converted
	expectedType interface{}, // nil pointer of desired type
	args ...argmapper.Arg, // list of argmapper arguments
) (interface{}, error) {
	args = append(args,
		argmapper.ConverterFunc(b.Mappers...),
		argmapper.Typed(b.Internal()),
		argmapper.Typed(b.Logger),
	)

	return dynamic.Map(resultValue, expectedType, args...)
}

// Set the cache to be used
func (b *Base) SetCache(c cacher.Cache) {
	b.Cache = c
}

// Set seed values. These values are will be automatically
// addeed to all dynamic calls.
func (b *BaseClient) Seed(
	args *core.Seeds, // typed and named values to store
) error {
	if b.Client == nil {
		b.Logger.Trace("plugin does not implement seeder interface")
		return nil
	}

	cb := func(d *vagrant_plugin_sdk.Args_Seeds) error {
		_, err := b.Client.(SeederClient).Seed(b.Ctx, d)
		return err
	}

	_, err := b.CallDynamicFunc(cb, false,
		argmapper.Typed(b.Ctx),
		argmapper.Typed(args),
	)

	return err
}

// Returns the collection of stored seed values
func (b *BaseClient) Seeds() (*core.Seeds, error) {
	if b.Client == nil {
		b.Logger.Trace("plugin does not implement seeder interface")
		return core.NewSeeds(), nil
	}

	r, err := b.Client.(SeederClient).Seeds(b.Ctx, &emptypb.Empty{})
	if err != nil {
		b.Logger.Error("failed to get seed values",
			"error", err,
		)

		return nil, err
	}

	s, err := b.Map(r, (**core.Seeds)(nil), argmapper.Typed(b.Ctx))
	if err != nil {
		b.Logger.Error("failed to convert seeds value response to proper type",
			"value", r,
			"error", err,
		)

		return nil, err
	}

	return s.(*core.Seeds), nil
}

// Set the plugin name
func (b *BaseClient) SetPluginName(name string) (err error) {
	if b.Client == nil {
		return
	}
	if c, ok := b.Client.(NamedPluginClient); ok {
		_, err = c.SetPluginName(
			b.Ctx, &vagrant_plugin_sdk.PluginInfo_Name{Name: name},
		)
		return
	}
	return errors.New("plugin does not support naming")
}

// Returns the name of the plugin
func (b *BaseClient) PluginName() (name string, err error) {
	if b.Client == nil {
		return
	}
	if c, ok := b.Client.(NamedPluginClient); ok {
		pluginName, err := c.PluginName(
			b.Ctx, &emptypb.Empty{},
		)
		if err != nil {
			return "", err
		}
		return pluginName.Name, nil

	}
	return "", errors.New("plugin does not support naming")
}

// Sets the parent component
func (b *BaseClient) SetParentComponent(c interface{}) {
	b.parentComponent = c
}

// Returns parent component
func (b *BaseClient) GetParentComponent() interface{} {
	return b.parentComponent
}

// Add extra mapper functions
func (b *BaseClient) AppendMappers(mappers ...*argmapper.Func) {
	b.Mappers = append(b.Mappers, mappers...)
}

// Generates a new context that is a combination of the
// given context and the client's context. Using the
// generated context helps to ensure things like customized
// metadata is included within client requests.
func (b *BaseClient) GenerateContext(ctx context.Context) (context.Context, context.CancelFunc) {
	return joincontext.Join(ctx, b.Ctx)
}

// Close the client and perform any required cleanup
func (b *BaseClient) Close() error {
	return b.Cleanup.Close()
}

// Used internally to extract broker
func (b *BaseClient) GRPCBroker() *plugin.GRPCBroker {
	return b.Broker
}

// Sets a direct addr which can be connected
// to when passing this client over proto.
func (b *BaseClient) SetAddr(t net.Addr) {
	b.addr = t
}

// Provides the direct addr being used
// by this client.
func (b *BaseClient) Addr() net.Addr {
	return b.addr
}

// This is here for internal usage on plugin setup
// to provide extra information to ruby Based plugins
func (b *BaseClient) SetRequestMetadata(key, value string) {
	md, ok := metadata.FromOutgoingContext(b.Ctx)
	if !ok {
		md = metadata.New(map[string]string{})
	}
	md[key] = []string{value}
	b.Ctx = metadata.NewOutgoingContext(b.Ctx, md)
	b.Logger.Trace("new metadata has been set for outgoing requests",
		"key", key, "value", value)
}

// Generate a function from a provided spec
func (b *BaseClient) GenerateFunc(
	spec *vagrant_plugin_sdk.FuncSpec, // spec for the function
	cbFn interface{}, // callback function
	args ...argmapper.Arg, // any extra argmapper args
) interface{} {
	return funcspec.Func(spec, cbFn, append(args,
		argmapper.Typed(b.Internal()))...,
	)
}

// Calls the function provided and converts the
// result to an expected type. If no type conversion
// is required, a `false` value for the expectedType
// will return the raw interface return value. Automatically
// provided args include hclog.Logger and pluginargs.Internal
// typed arguments, registered mappers, and a custom logger
// for argmapper.
//
// NOTE: Provide a `false` value for expectedType if no
// type conversion is required.
func (b *BaseClient) CallDynamicFunc(
	f interface{}, // function to call
	expectedType interface{}, // nil pointer of expected return type
	callArgs ...argmapper.Arg, // any extra argmapper arguments to include
) (interface{}, error) {
	internal := b.Internal()

	if b.Client != nil {
		s, err := b.Seeds()
		if err != nil {
			b.Logger.Error("failed to fetch dynamic seed values",
				"error", err,
			)

			return nil, err
		}

		for _, v := range s.Typed {
			// If the value is an Any value, unpack it before
			// adding the value
			if a, ok := v.(*anypb.Any); ok {
				val, err := a.UnmarshalNew()
				if err != nil {
					b.Logger.Info("failed to unmarshal Any type seed value",
						"value", a,
						"error", err,
					)

					return nil, err
				}

				b.Logger.Trace("seeding typed value into dynamic call",
					"type", hclog.Fmt("%T", val),
				)

				callArgs = append(callArgs, argmapper.Typed(val))
			} else {
				b.Logger.Trace("seeding typed value into dynamic call",
					"type", hclog.Fmt("%T", v),
				)

				callArgs = append(callArgs, argmapper.Typed(v))
			}
		}

		for k := range s.Named {
			v := s.Named[k]
			// If the value is an Any value, unpack it before
			// adding the value
			if a, ok := v.(*anypb.Any); ok {
				val, err := a.UnmarshalNew()
				if err != nil {
					b.Logger.Info("failed to unmarshal Any type seed value",
						"value", a,
						"error", err,
					)

					return nil, err
				}

				b.Logger.Trace("seeding named value into dynamic call",
					"name", k,
					"type", hclog.Fmt("%T", val),
				)

				callArgs = append(callArgs, argmapper.Named(k, val))
			} else {
				b.Logger.Trace("seeding named value into dynamic call",
					"name", k,
					"type", hclog.Fmt("%T", v),
				)

				callArgs = append(callArgs, argmapper.Named(k, v))
			}
		}
	}

	callArgs = append(callArgs,
		argmapper.Typed(internal),
		argmapper.Typed(b.Logger),
	)

	return dynamic.CallFunc(f, expectedType, b.Mappers, callArgs...)
}

// Calls the function provided and converts the
// result to an expected type. If no type conversion
// is required, a `false` value for the expectedType
// will return the raw interface return value. Automatically
// provided args include hclog.Logger and pluginargs.Internal
// typed arguments, registered mappers, and a custom logger
// for argmapper.
//
// NOTE: Provide a `false` value for expectedType if no
// type conversion is required.
func (b *BaseServer) CallDynamicFunc(
	f interface{}, // function to call
	expectedType interface{}, // nil pointer of expected return type
	args funcspec.Args, // funspec defined arguments
	callArgs ...argmapper.Arg, // any extra argmapper arguments to include
) (interface{}, error) {
	internal := b.Internal()

	// Decode our *anypb.Any values.
	for _, arg := range args {
		anyVal := arg.Value

		_, v, err := dynamic.DecodeAny(anyVal)
		if err != nil {
			return nil, err
		}

		callArgs = append(callArgs,
			argmapper.NamedSubtype(arg.Name, v, arg.Type),
		)
	}
	callArgs = append(callArgs,
		argmapper.Typed(internal),
		argmapper.Typed(b.Logger),
	)

	return dynamic.CallFunc(f, expectedType, b.Mappers, callArgs...)
}

// Generate a funcspec Based on the provided function
func (b *BaseServer) GenerateSpec(
	fn interface{}, // function to generate funcspec
	args ...argmapper.Arg, // optional argmapper args
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if f, ok := fn.(*dynamic.SpecAndFunc); ok {
		return f.Spec, nil
	}
	f, err := funcspec.Spec(fn,
		append(
			args,
			argmapper.ConverterFunc(b.Mappers...),
			argmapper.Typed(
				b.Internal(),
				b.Logger,
			),
		)...,
	)
	if err != nil {
		return f, err
	}
	return f, err
}

// Store seed values
func (b *BaseServer) Seed(
	ctx context.Context,
	seeds *vagrant_plugin_sdk.Args_Seeds,
) (*emptypb.Empty, error) {
	if b.impl == nil {
		b.Logger.Trace("plugin does not implement seeder interface")
		return &emptypb.Empty{}, nil
	}

	if !b.IsWrapped() {
		b.seedValues = seeds
		return &emptypb.Empty{}, nil
	}

	seeder, ok := b.impl.(core.Seeder)
	if !ok {
		b.Logger.Error("plugin implementation does not provide core.Seeder",
			"impl", b.impl,
		)

		return nil, fmt.Errorf("implementation does not support value seeds")
	}

	v, err := dynamic.Map(seeds, (**core.Seeds)(nil),
		argmapper.Typed(ctx, b.Internal(), b.Logger),
		argmapper.ConverterFunc(b.Mappers...),
	)

	if err != nil {
		b.Logger.Error("failed to store seed values",
			"error", err,
		)

		return nil, err
	}

	err = seeder.Seed(v.(*core.Seeds))

	if err != nil {
		b.Logger.Error("failed to store seed values",
			"error", err,
		)
	}
	return &emptypb.Empty{}, err
}

// Returns collection of stored seed values
func (b *BaseServer) Seeds(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Args_Seeds, error) {
	if b.impl == nil {
		b.Logger.Trace("plugin does not implement seeder interface")
		return &vagrant_plugin_sdk.Args_Seeds{}, nil
	}

	if !b.IsWrapped() {
		return b.seedValues, nil
	}

	seeder, ok := b.impl.(core.Seeder)
	if !ok {
		b.Logger.Error("plugin implementation does not provide core.Seeder",
			"impl", b.impl,
		)

		return nil, fmt.Errorf("implementation does not support value seeds")
	}

	s, err := seeder.Seeds()
	if err != nil {
		b.Logger.Error("failed to fetch seed values",
			"error", err,
		)

		return nil, err
	}

	r, err := dynamic.Map(s,
		(**vagrant_plugin_sdk.Args_Seeds)(nil),
		argmapper.Typed(ctx, b.Internal(), b.Logger),
		argmapper.ConverterFunc(b.Mappers...),
	)

	if err != nil {
		b.Logger.Error("failed to convert seed values into proto message",
			"values", s,
			"error", err,
		)

		return nil, err
	}

	return r.(*vagrant_plugin_sdk.Args_Seeds), nil
}

// Set the plugin name
func (b *BaseServer) SetPluginName(
	ctx context.Context,
	name *vagrant_plugin_sdk.PluginInfo_Name,
) (*emptypb.Empty, error) {
	if !b.IsWrapped() {
		b.name = name.Name
		return &emptypb.Empty{}, nil
	}

	if namedPlugin, ok := b.impl.(core.Named); ok {
		err := namedPlugin.SetPluginName(name.Name)
		if err != nil {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}

// Returns the plugin name
func (b *BaseServer) PluginName(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.PluginInfo_Name, error) {
	if !b.IsWrapped() {
		return &vagrant_plugin_sdk.PluginInfo_Name{Name: b.name}, nil
	}

	if namedPlugin, ok := b.impl.(core.Named); ok {
		name, err := namedPlugin.PluginName()
		if err != nil {
			return nil, err
		}
		return &vagrant_plugin_sdk.PluginInfo_Name{Name: name}, nil
	}
	return &vagrant_plugin_sdk.PluginInfo_Name{Name: ""}, nil
}

// Generate full mapper list. This will add our locally
// defined mappers to the given list, ensuring duplicates
// aren't added.
func mappers(base []*argmapper.Func) []*argmapper.Func {
	result := make([]*argmapper.Func, len(base))
	copy(result, base)
	for _, m := range MapperFns {
		found := false
		for _, e := range result {
			if m == e {
				found = true
				break
			}
		}
		if !found {
			result = append(result, m)
		}
	}

	return result
}

var (
	_ core.Named  = (*BaseClient)(nil)
	_ core.Seeder = (*BaseClient)(nil)
)
