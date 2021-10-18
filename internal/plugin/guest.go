package plugin

import (
	"context"

	"github.com/LK4D4/joincontext"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/docs"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/funcspec"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
)

// GuestPlugin implements plugin.Plugin (specifically GRPCPlugin) for
// the Guest component type.
type GuestPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl component.Guest // Impl is the concrete implementation
	*BasePlugin
}

func (p *GuestPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	bs := p.NewServer(broker)
	vagrant_plugin_sdk.RegisterGuestServiceServer(s, &guestServer{
		Impl:       p.Impl,
		BaseServer: bs,
		capabilityServer: &capabilityServer{
			BaseServer:     bs,
			CapabilityImpl: p.Impl,
			typ:            "guest",
		},
	})
	return nil
}

func (p *GuestPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	bc := p.NewClient(ctx, broker)
	client := vagrant_plugin_sdk.NewGuestServiceClient(c)
	return &guestClient{
		client:     client,
		BaseClient: bc,
		capabilityClient: &capabilityClient{
			client:     client,
			BaseClient: bc,
		},
	}, nil
}

// guestClient is an implementation of component.Guest over gRPC.
type guestClient struct {
	*BaseClient
	*capabilityClient
	client vagrant_plugin_sdk.GuestServiceClient
}

func (c *guestClient) GetCapabilityClient() *capabilityClient {
	return c.capabilityClient
}

func (c *guestClient) Config() (interface{}, error) {
	return configStructCall(c.Ctx, c.client)
}

func (c *guestClient) ConfigSet(v interface{}) error {
	return configureCall(c.Ctx, c.client, v)
}

func (c *guestClient) Documentation() (*docs.Documentation, error) {
	return documentationCall(c.Ctx, c.client)
}

func (c *guestClient) GuestDetectFunc() interface{} {
	spec, err := c.client.DetectSpec(c.Ctx, &empty.Empty{})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil
	cb := func(ctx context.Context, args funcspec.Args) (bool, error) {
		resp, err := c.client.Detect(ctx, &vagrant_plugin_sdk.FuncSpec_Args{Args: args})
		if err != nil {
			return false, err
		}
		return resp.Detected, nil
	}
	return c.GenerateFunc(spec, cb)
}

func (c *guestClient) Detect(t core.Target) (bool, error) {
	f := c.GuestDetectFunc()
	raw, err := c.CallDynamicFunc(f, (*bool)(nil),
		argmapper.Typed(c.Ctx),
		argmapper.Typed(t),
	)
	if err != nil {
		return false, err
	}

	return raw.(bool), nil
}

func (c *guestClient) ParentsFunc() interface{} {
	spec, err := c.client.ParentsSpec(c.Ctx, &empty.Empty{})
	if err != nil {
		return funcErr(err)
	}
	spec.Result = nil
	cb := func(ctx context.Context, args funcspec.Args) ([]string, error) {
		ctx, _ = joincontext.Join(c.Ctx, ctx)
		resp, err := c.client.Parents(ctx, &vagrant_plugin_sdk.FuncSpec_Args{Args: args})
		if err != nil {
			return nil, err
		}
		return resp.Parents, nil
	}

	return c.GenerateFunc(spec, cb)
}

func (c *guestClient) Parents() ([]string, error) {
	f := c.ParentsFunc()
	raw, err := c.CallDynamicFunc(f, (*[]string)(nil),
		argmapper.Typed(c.Ctx),
	)
	if err != nil {
		return nil, err
	}

	return raw.([]string), nil
}

// guestServer is a gRPC server that the client talks to and calls a
// real implementation of the component.
type guestServer struct {
	*BaseServer
	*capabilityServer

	Impl component.Guest
}

func (s *guestServer) ConfigStruct(
	ctx context.Context,
	empty *empty.Empty,
) (*vagrant_plugin_sdk.Config_StructResp, error) {
	return configStruct(s.Impl)
}

func (s *guestServer) Configure(
	ctx context.Context,
	req *vagrant_plugin_sdk.Config_ConfigureRequest,
) (*empty.Empty, error) {
	return configure(s.Impl, req)
}

func (s *guestServer) Documentation(
	ctx context.Context,
	empty *empty.Empty,
) (*vagrant_plugin_sdk.Config_Documentation, error) {
	return documentation(s.Impl)
}

func (s *guestServer) DetectSpec(
	ctx context.Context,
	args *empty.Empty,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, "guest"); err != nil {
		return nil, err
	}

	return s.GenerateSpec(s.Impl.GuestDetectFunc())
}

func (s *guestServer) Detect(
	ctx context.Context,
	args *vagrant_plugin_sdk.FuncSpec_Args,
) (*vagrant_plugin_sdk.Platform_DetectResp, error) {
	raw, err := s.CallDynamicFunc(s.Impl.GuestDetectFunc(), (*bool)(nil), args.Args,
		argmapper.Typed(ctx),
	)

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Platform_DetectResp{Detected: raw.(bool)}, nil
}

func (s *guestServer) ParentsSpec(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.FuncSpec, error) {
	if err := isImplemented(s, s.typ); err != nil {
		return nil, err
	}

	return s.GenerateSpec(s.Impl.ParentsFunc())
}

func (s *guestServer) Parents(
	ctx context.Context,
	args *vagrant_plugin_sdk.FuncSpec_Args,
) (*vagrant_plugin_sdk.Platform_ParentsResp, error) {
	raw, err := s.CallDynamicFunc(s.Impl.ParentsFunc(), (*[]string)(nil),
		args.Args, argmapper.Typed(ctx))

	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Platform_ParentsResp{
		Parents: raw.([]string)}, nil
}

var (
	_ plugin.Plugin                         = (*GuestPlugin)(nil)
	_ plugin.GRPCPlugin                     = (*GuestPlugin)(nil)
	_ vagrant_plugin_sdk.GuestServiceServer = (*guestServer)(nil)
	_ component.Guest                       = (*guestClient)(nil)
	_ core.Guest                            = (*guestClient)(nil)
	_ capabilityParent                      = (*hostClient)(nil)
)
