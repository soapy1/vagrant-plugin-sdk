package core

import (
	"context"

	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/hashicorp/vagrant-plugin-sdk/component"
	"github.com/hashicorp/vagrant-plugin-sdk/core"
	vplugin "github.com/hashicorp/vagrant-plugin-sdk/internal/plugin"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
)

// VagrantfilePlugin is just a GRPC client for a vagrantfile
type VagrantfilePlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl core.Vagrantfile
	*vplugin.BasePlugin
}

// Implements plugin.GRPCPlugin
func (p *VagrantfilePlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &vagrantfileClient{
		client:     vagrant_plugin_sdk.NewVagrantfileServiceClient(c),
		BaseClient: p.NewClient(ctx, broker, nil),
	}, nil
}

func (p *VagrantfilePlugin) GRPCServer(
	broker *plugin.GRPCBroker,
	s *grpc.Server,
) error {
	vagrant_plugin_sdk.RegisterVagrantfileServiceServer(s, &vagrantfileServer{
		Impl:       p.Impl,
		BaseServer: p.NewServer(broker, nil),
	})
	return nil
}

type vagrantfileClient struct {
	*vplugin.BaseClient

	client vagrant_plugin_sdk.VagrantfileServiceClient
}

type vagrantfileServer struct {
	*vplugin.BaseServer

	Impl core.Vagrantfile
	vagrant_plugin_sdk.UnimplementedVagrantfileServiceServer
}

func (v *vagrantfileClient) Target(
	name, provider string,
) (machine core.Target, err error) {
	resp, err := v.client.Target(v.Ctx,
		&vagrant_plugin_sdk.Vagrantfile_TargetRequest{
			Name:     name,
			Provider: provider,
		},
	)
	if err != nil {
		return nil, err
	}

	raw, err := v.Map(resp, (*core.Target)(nil), argmapper.Typed(v.Ctx))
	if err != nil {
		return nil, err
	}

	return raw.(core.Target), nil
}

func (v *vagrantfileClient) TargetConfig(
	name, provider string,
	validateProvider bool,
) (config core.Vagrantfile, err error) {
	resp, err := v.client.TargetConfig(v.Ctx,
		&vagrant_plugin_sdk.Vagrantfile_TargetConfigRequest{
			Name:             name,
			Provider:         provider,
			ValidateProvider: validateProvider,
		},
	)
	if err != nil {
		return nil, err
	}

	raw, err := v.Map(resp, (*core.Vagrantfile)(nil), argmapper.Typed(v.Ctx))
	if err != nil {
		return nil, err
	}

	return raw.(core.Vagrantfile), nil
}

func (v *vagrantfileClient) TargetNames() ([]string, error) {
	resp, err := v.client.TargetNames(v.Ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return resp.Names, nil
}

func (v *vagrantfileClient) PrimaryTargetName() (string, error) {
	resp, err := v.client.PrimaryTargetName(v.Ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}

	return resp.Name, nil
}

func (v *vagrantfileClient) GetConfig(namespace string) (*component.ConfigData, error) {
	resp, err := v.client.GetConfig(v.Ctx,
		&vagrant_plugin_sdk.Vagrantfile_NamespaceRequest{
			Namespace: namespace,
		},
	)
	if err != nil {
		return nil, err
	}
	raw, err := v.Map(resp, (**component.ConfigData)(nil), argmapper.Typed(v.Ctx))
	if err != nil {
		return nil, err
	}

	return raw.(*component.ConfigData), nil
}

func (v *vagrantfileClient) GetValue(path ...string) (interface{}, error) {
	resp, err := v.client.GetValue(v.Ctx,
		&vagrant_plugin_sdk.Vagrantfile_ValueRequest{
			Path: path,
		},
	)
	if err != nil {
		return nil, err
	}
	raw, err := v.Map(resp, (*component.Direct)(nil), argmapper.Typed(v.Ctx))
	if err != nil {
		return nil, err
	}

	return raw.(*component.Direct).Arguments[0], nil
}

// Server

func (v *vagrantfileServer) Target(
	ctx context.Context,
	req *vagrant_plugin_sdk.Vagrantfile_TargetRequest,
) (*vagrant_plugin_sdk.Args_Target, error) {
	t, err := v.Impl.Target(req.Name, req.Provider)
	if err != nil {
		v.Logger.Error("failed to get target from implementation",
			"error", err,
		)

		return nil, err
	}

	raw, err := v.Map(t, (**vagrant_plugin_sdk.Args_Target)(nil), argmapper.Typed(ctx))
	if err != nil {
		return nil, err
	}

	return raw.(*vagrant_plugin_sdk.Args_Target), nil
}

func (v *vagrantfileServer) TargetConfig(
	ctx context.Context,
	req *vagrant_plugin_sdk.Vagrantfile_TargetConfigRequest,
) (*vagrant_plugin_sdk.Args_Vagrantfile, error) {
	vf, err := v.Impl.TargetConfig(req.Name, req.Provider, req.ValidateProvider)
	if err != nil {
		v.Logger.Error("failed to get target config from implementation",
			"error", err,
		)
		return nil, err
	}

	raw, err := v.Map(vf, (**vagrant_plugin_sdk.Args_Vagrantfile)(nil), argmapper.Typed(ctx))
	if err != nil {
		v.Logger.Error("failed to map vagrantfile for config response",
			"error", err,
		)
		return nil, err
	}

	return raw.(*vagrant_plugin_sdk.Args_Vagrantfile), nil
}

func (v *vagrantfileServer) TargetNames(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Vagrantfile_TargetNamesResponse, error) {
	n, err := v.Impl.TargetNames()
	if err != nil {
		v.Logger.Error("failed to get target names from implementation",
			"error", err,
		)
		return nil, err
	}
	return &vagrant_plugin_sdk.Vagrantfile_TargetNamesResponse{
		Names: n,
	}, nil
}

func (v *vagrantfileServer) PrimaryTargetName(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Vagrantfile_PrimaryTargetNameResponse, error) {
	n, err := v.Impl.PrimaryTargetName()
	if err != nil {
		v.Logger.Error("failed to get primary target name from implementation",
			"error", err,
		)
		return nil, err
	}

	return &vagrant_plugin_sdk.Vagrantfile_PrimaryTargetNameResponse{
		Name: n,
	}, nil
}

func (v *vagrantfileServer) GetConfig(
	ctx context.Context,
	req *vagrant_plugin_sdk.Vagrantfile_NamespaceRequest,
) (*vagrant_plugin_sdk.Args_ConfigData, error) {
	c, err := v.Impl.GetConfig(req.Namespace)
	if err != nil {
		v.Logger.Error("failed to get config from implementation",
			"namespace", req.Namespace,
			"error", err,
		)
		return nil, err
	}
	raw, err := v.Map(c, (**vagrant_plugin_sdk.Args_ConfigData)(nil), argmapper.Typed(ctx))
	if err != nil {
		v.Logger.Error("failed to map config data",
			"error", err,
		)
		return nil, err
	}
	return raw.(*vagrant_plugin_sdk.Args_ConfigData), nil
}

func (v *vagrantfileServer) GetValue(
	ctx context.Context,
	req *vagrant_plugin_sdk.Vagrantfile_ValueRequest,
) (*vagrant_plugin_sdk.Args_Direct, error) {
	var val interface{}

	val, err := v.Impl.GetValue(req.Path...)
	if err != nil {
		v.Logger.Error("failed to get config value from implementation",
			"path", req.Path,
			"error", err,
		)
		return nil, err
	}

	if sh, ok := val.(map[string]interface{}); ok {
		ih := make(map[interface{}]interface{}, len(sh))
		for k, v := range sh {
			ih[k] = v
		}
		val = ih
	}

	v.Logger.Info("got value from vagrantfile config for conversion",
		"value", val,
	)

	raw, err := v.Map(
		&component.Direct{Arguments: []interface{}{val}},
		(*proto.Message)(nil),
		argmapper.Typed(ctx),
	)

	if err != nil {
		return nil, err
	}

	v.Logger.Info("got value and converting to any value",
		"value", raw,
	)

	return raw.(*vagrant_plugin_sdk.Args_Direct), nil
}

var (
	_ plugin.Plugin     = (*VagrantfilePlugin)(nil)
	_ plugin.GRPCPlugin = (*VagrantfilePlugin)(nil)
	_ core.Vagrantfile  = (*vagrantfileClient)(nil)
)
