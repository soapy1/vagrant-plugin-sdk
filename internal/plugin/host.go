package plugin

import (
	"context"

	"github.com/LK4D4/joincontext"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/docs"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/funcspec"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
)

// HostPlugin implements plugin.Plugin (specifically GRPCPlugin) for
// the Host component type.
type HostPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl    component.Host    // Impl is the concrete implementation
	Mappers []*argmapper.Func // Mappers
	Logger  hclog.Logger      // Logger
}

func (p *HostPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	vagrant_plugin_sdk.RegisterHostServiceServer(s, &hostServer{
		Impl: p.Impl,
		baseServer: &baseServer{
			base: &base{
				Mappers: p.Mappers,
				Logger:  p.Logger,
				Broker:  broker,
			},
		},
	})
	return nil
}

func (p *HostPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &hostClient{
		client: vagrant_plugin_sdk.NewHostServiceClient(c),
		baseClient: &baseClient{
			ctx: context.Background(),
			base: &base{
				Mappers: p.Mappers,
				Logger:  p.Logger,
				Broker:  broker,
			},
		},
		capabilityClient: &capabilityClient{
			client: vagrant_plugin_sdk.NewHostServiceClient(c),
		},
	}, nil
}

type hostClient struct {
	*baseClient
	*capabilityClient
	client vagrant_plugin_sdk.HostServiceClient
}

func (c *hostClient) Config() (interface{}, error) {
	return configStructCall(c.ctx, c.client)
}

func (c *hostClient) ConfigSet(v interface{}) error {
	return configureCall(c.ctx, c.client, v)
}

func (c *hostClient) Documentation() (*docs.Documentation, error) {
	return documentationCall(c.ctx, c.client)
}

func (c *hostClient) DetectFunc() interface{} {
	spec, err := c.client.DetectSpec(c.ctx, &empty.Empty{})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil
	cb := func(ctx context.Context, args funcspec.Args) (bool, error) {
		ctx, _ = joincontext.Join(c.ctx, ctx)
		resp, err := c.client.Detect(ctx, &vagrant_plugin_sdk.FuncSpec_Args{Args: args})
		if err != nil {
			return false, err
		}
		return resp.Detected, nil
	}

	return c.generateFunc(spec, cb)
}

func (c *hostClient) Detect() (bool, error) {
	f := c.DetectFunc()
	raw, err := c.callDynamicFunc(f, (*bool)(nil),
		argmapper.Typed(c.ctx),
	)
	if err != nil {
		return false, err
	}

	return raw.(bool), nil
}

func (c *hostClient) ParentsFunc() interface{} {
	spec, err := c.client.ParentsSpec(c.ctx, &empty.Empty{})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil
	cb := func(ctx context.Context, args funcspec.Args) ([]string, error) {
		ctx, _ = joincontext.Join(c.ctx, ctx)
		resp, err := c.client.Parents(ctx, &vagrant_plugin_sdk.FuncSpec_Args{Args: args})
		if err != nil {
			return nil, err
		}
		return resp.Parents, nil
	}

	return c.generateFunc(spec, cb)
}

func (c *hostClient) Parents() ([]string, error) {
	f := c.ParentsFunc()
	raw, err := c.callDynamicFunc(f, (*[]string)(nil),
		argmapper.Typed(c.ctx),
	)
	if err != nil {
		return nil, err
	}

	return raw.([]string), nil
}

func (c *hostClient) HasCapabilityFunc() interface{} {
	spec, err := c.client.HasCapabilitySpec(c.ctx, &empty.Empty{})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil

	cb := func(ctx context.Context, args funcspec.Args) (bool, error) {
		ctx, _ = joincontext.Join(c.ctx, ctx)
		resp, err := c.client.HasCapability(ctx, &vagrant_plugin_sdk.FuncSpec_Args{Args: args})

		if err != nil {
			return false, err
		}
		return resp.HasCapability, nil
	}
	return c.generateFunc(spec, cb)
}

func (c *hostClient) HasCapability(name string) (bool, error) {
	f := c.HasCapabilityFunc()
	n := &component.NamedCapability{Capability: name}
	raw, err := c.callDynamicFunc(f, (*bool)(nil),
		argmapper.Typed(n),
		argmapper.Typed(c.ctx),
	)
	if err != nil {
		return false, err
	}

	return raw.(bool), nil
}

func (c *hostClient) CapabilityFunc(name string) interface{} {
	spec, err := c.client.CapabilitySpec(c.ctx,
		&vagrant_plugin_sdk.Host_Capability_NamedRequest{Name: name})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil
	cb := func(ctx context.Context, args funcspec.Args) (interface{}, error) {
		ctx, _ = joincontext.Join(c.ctx, ctx)
		resp, err := c.client.Capability(ctx,
			&vagrant_plugin_sdk.Host_Capability_NamedRequest{
				FuncArgs: &vagrant_plugin_sdk.FuncSpec_Args{Args: args},
				Name:     name,
			},
		)

		if err != nil {
			return nil, err
		}
		return resp.Result, nil
	}
	return c.generateFunc(spec, cb)
}

func (c *hostClient) Capability(name string, args ...interface{}) (interface{}, error) {
	f := c.CapabilityFunc(name)
	raw, err := c.callDynamicFunc(f, false,
		argmapper.Typed(args...),
		argmapper.Typed(c.ctx),
	)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

type hostServer struct {
	*baseServer

	Impl component.Host
	vagrant_plugin_sdk.UnimplementedHostServiceServer
}

func (s *hostServer) ConfigStruct(
	ctx context.Context,
	empty *empty.Empty,
) (*vagrant_plugin_sdk.Config_StructResp, error) {
	return configStruct(s.Impl)
}

func (s *hostServer) Configure(
	ctx context.Context,
	req *vagrant_plugin_sdk.Config_ConfigureRequest,
) (*empty.Empty, error) {
	return configure(s.Impl, req)
}

func (s *hostServer) Documentation(
	ctx context.Context,
	empty *empty.Empty,
) (*vagrant_plugin_sdk.Config_Documentation, error) {
	return documentation(s.Impl)
}

func (s *hostServer) HasCapabilitySpec(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, "host"); err != nil {
		return nil, err
	}

	return s.generateSpec(s.Impl.HasCapabilityFunc())
}

func (s *hostServer) HasCapability(
	ctx context.Context,
	args *vagrant_plugin_sdk.FuncSpec_Args,
) (*vagrant_plugin_sdk.Host_Capability_CheckResp, error) {
	raw, err := s.callDynamicFunc(s.Impl.HasCapabilityFunc(), (*bool)(nil),
		args.Args, argmapper.Typed(ctx))

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Host_Capability_CheckResp{
		HasCapability: raw.(bool)}, nil
}

func (s *hostServer) CapabilitySpec(
	ctx context.Context,
	req *vagrant_plugin_sdk.Host_Capability_NamedRequest,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, "host"); err != nil {
		return nil, err
	}

	return s.generateSpec(s.Impl.CapabilityFunc(req.Name))
}

func (s *hostServer) Capability(
	ctx context.Context,
	args *vagrant_plugin_sdk.Host_Capability_NamedRequest,
) (*vagrant_plugin_sdk.Host_Capability_Resp, error) {
	_, err := s.callDynamicFunc(s.Impl.CapabilityFunc(args.Name), false,
		args.FuncArgs.Args, argmapper.Typed(ctx))

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Host_Capability_Resp{}, nil
}

func (s *hostServer) DetectSpec(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, "host"); err != nil {
		return nil, err
	}

	if err := isImplemented(s.Impl, "host"); err != nil {
		return nil, err
	}

	return s.generateSpec(s.Impl.DetectFunc())
}

func (s *hostServer) Detect(
	ctx context.Context,
	args *vagrant_plugin_sdk.FuncSpec_Args,
) (*vagrant_plugin_sdk.Host_DetectResp, error) {
	raw, err := s.callDynamicFunc(s.Impl.DetectFunc(), (*bool)(nil),
		args.Args, argmapper.Typed(ctx))

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Host_DetectResp{
		Detected: raw.(bool)}, nil
}

func (s *hostServer) ParentsSpec(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, "host"); err != nil {
		return nil, err
	}

	return s.generateSpec(s.Impl.ParentsFunc())
}

func (s *hostServer) Parents(
	ctx context.Context,
	args *vagrant_plugin_sdk.FuncSpec_Args,
) (*vagrant_plugin_sdk.Host_ParentsResp, error) {
	raw, err := s.callDynamicFunc(s.Impl.ParentsFunc(), (*[]string)(nil),
		args.Args, argmapper.Typed(ctx))

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Host_ParentsResp{
		Parents: raw.([]string)}, nil
}

var (
	_ plugin.Plugin                        = (*HostPlugin)(nil)
	_ plugin.GRPCPlugin                    = (*HostPlugin)(nil)
	_ vagrant_plugin_sdk.HostServiceServer = (*hostServer)(nil)
	_ component.Host                       = (*hostClient)(nil)
	_ core.Host                            = (*hostClient)(nil)
)
