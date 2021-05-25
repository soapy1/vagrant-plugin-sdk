package core

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/datadir"
	"github.com/hashicorp/vagrant-plugin-sdk/internal/pluginargs"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// ProjectPlugin is just a GRPC client for a project
type ProjectPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Mappers []*argmapper.Func // Mappers
	Logger  hclog.Logger      // Logger
	Impl    core.Project
}

// Implements plugin.GRPCPlugin
func (p *ProjectPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &projectClient{
		client: vagrant_plugin_sdk.NewProjectServiceClient(c),
		ctx:    ctx,
		base: &base{
			Mappers: p.Mappers,
			Logger:  p.Logger,
			Broker:  broker,
			Cleanup: &pluginargs.Cleanup{},
		},
	}, nil
}

func (p *ProjectPlugin) GRPCServer(
	broker *plugin.GRPCBroker,
	s *grpc.Server,
) error {
	vagrant_plugin_sdk.RegisterProjectServiceServer(s, &projectServer{
		Impl: p.Impl,
		base: &base{
			Mappers: p.Mappers,
			Logger:  p.Logger,
			Broker:  broker,
			Cleanup: &pluginargs.Cleanup{},
		},
	})
	return nil
}

type projectClient struct {
	*base

	ctx    context.Context
	client vagrant_plugin_sdk.ProjectServiceClient
}

type projectServer struct {
	*base

	Impl core.Project
	vagrant_plugin_sdk.UnimplementedProjectServiceServer
}

func (p *projectClient) CWD() (path string, err error) {
	r, err := p.client.CWD(p.ctx, &emptypb.Empty{})
	if err == nil {
		path = r.Path
	}

	return
}

func (p *projectClient) Target(name string) (t core.Target, err error) {
	r, err := p.client.Target(p.ctx, &vagrant_plugin_sdk.Project_TargetRequest{
		Name: name,
	})
	if err != nil {
		return
	}

	result, err := p.Map(r, (*core.Target)(nil),
		argmapper.Typed(p.ctx))
	if err == nil {
		t = result.(core.Target)
	}
	return
}

func (p *projectClient) TargetNames() (names []string, err error) {
	r, err := p.client.TargetNames(p.ctx, &empty.Empty{})
	if err == nil {
		names = r.Names
	}

	return
}

func (p *projectClient) TargetIds() (ids []string, err error) {
	r, err := p.client.TargetIds(p.ctx, &empty.Empty{})
	if err == nil {
		ids = r.Ids
	}

	return
}

func (p *projectClient) DataDir() (dir *datadir.Project, err error) {
	r, err := p.client.DataDir(p.ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	result, err := p.Map(r, (**datadir.Project)(nil))
	if err == nil {
		dir = result.(*datadir.Project)
	}

	return
}

func (p *projectClient) VagrantfileName() (name string, err error) {
	r, err := p.client.VagrantfileName(p.ctx, &emptypb.Empty{})
	if err == nil {
		name = r.Name
	}

	return
}

func (p *projectClient) UI() (ui terminal.UI, err error) {
	r, err := p.client.UI(p.ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	result, err := p.Map(r, (*terminal.UI)(nil),
		argmapper.Typed(p.ctx))
	if err == nil {
		ui = result.(terminal.UI)
	}

	return
}

func (p *projectClient) Home() (path string, err error) {
	r, err := p.client.Home(p.ctx, &emptypb.Empty{})
	if err == nil {
		path = r.Path
	}

	return
}
func (p *projectClient) LocalData() (path string, err error) {
	r, err := p.client.LocalData(p.ctx, &emptypb.Empty{})
	if err == nil {
		path = r.Path
	}

	return
}

func (p *projectClient) Tmp() (path string, err error) {
	r, err := p.client.Tmp(p.ctx, &emptypb.Empty{})
	if err == nil {
		path = r.Path
	}

	return
}

func (p *projectClient) DefaultPrivateKey() (path string, err error) {
	r, err := p.client.DefaultPrivateKey(p.ctx, &emptypb.Empty{})
	if err == nil {
		path = r.Key
	}

	return
}

func (p *projectClient) MachineNames() (names []string, err error) {
	r, err := p.client.MachineNames(p.ctx, &emptypb.Empty{})
	if err == nil {
		names = r.Names
	}

	return
}

func (p *projectClient) Host() (h core.Host, err error) {
	r, err := p.client.Host(p.ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	result, err := p.Map(r, (*core.Host)(nil))
	if err == nil {
		h = result.(core.Host)
	}

	return
}

func (p *projectServer) CWD(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_CwdResponse, error) {
	c, err := p.Impl.CWD()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_CwdResponse{
		Path: c,
	}, nil
}

func (p *projectServer) DataDir(
	ctx context.Context,
	_ *empty.Empty,
) (r *vagrant_plugin_sdk.Args_DataDir_Project, err error) {
	d, err := p.Impl.DataDir()
	if err != nil {
		return
	}
	result, err := p.Map(d, (**vagrant_plugin_sdk.Args_DataDir_Project)(nil))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_DataDir_Project)
	}

	return
}

func (p *projectServer) VagrantfileName(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_VagrantfileNameResponse, error) {
	name, err := p.Impl.VagrantfileName()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_VagrantfileNameResponse{
		Name: name,
	}, nil
}

func (p *projectServer) UI(
	ctx context.Context,
	_ *empty.Empty,
) (r *vagrant_plugin_sdk.Args_TerminalUI, err error) {
	d, err := p.Impl.UI()
	if err != nil {
		return
	}

	result, err := p.Map(d, (**vagrant_plugin_sdk.Args_TerminalUI)(nil),
		argmapper.Typed(ctx))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_TerminalUI)
	}

	return
}

func (p *projectServer) Home(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_HomeResponse, error) {
	path, err := p.Impl.Home()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_HomeResponse{
		Path: path,
	}, nil
}

func (p *projectServer) LocalData(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_LocalDataResponse, error) {
	path, err := p.Impl.LocalData()
	if err != nil {
		return nil, err
	}
	return &vagrant_plugin_sdk.Project_LocalDataResponse{
		Path: path,
	}, nil
}

func (p *projectServer) Tmp(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_TmpResponse, error) {
	path, err := p.Impl.Tmp()
	if err != nil {
		return nil, err
	}
	return &vagrant_plugin_sdk.Project_TmpResponse{
		Path: path,
	}, nil
}

func (p *projectServer) DefaultPrivateKey(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_DefaultPrivateKeyResponse, error) {
	key, err := p.Impl.DefaultPrivateKey()
	p.Logger.Warn("private key on project server", "key", key)
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_DefaultPrivateKeyResponse{
		Key: key,
	}, nil
}

func (p *projectServer) Host(
	ctx context.Context,
	_ *empty.Empty,
) (r *vagrant_plugin_sdk.Args_Host, err error) {
	d, err := p.Impl.Host()
	if err != nil {
		return
	}

	result, err := p.Map(d, (**vagrant_plugin_sdk.Args_Host)(nil))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_Host)
	}

	return
}

func (p *projectServer) MachineNames(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_MachineNamesResponse, error) {
	names, err := p.Impl.MachineNames()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_MachineNamesResponse{
		Names: names,
	}, nil
}

func (p *projectServer) Target(
	ctx context.Context,
	in *vagrant_plugin_sdk.Project_TargetRequest,
) (r *vagrant_plugin_sdk.Args_Target, err error) {
	d, err := p.Impl.Target(in.Name)
	if err != nil {
		return
	}

	result, err := p.Map(d, (**vagrant_plugin_sdk.Args_Target)(nil))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_Target)
	}

	return
}

func (p *projectServer) TargetNames(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_TargetNamesResponse, error) {
	n, err := p.Impl.TargetNames()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_TargetNamesResponse{
		Names: n}, nil
}

func (p *projectServer) TargetIds(
	ctx context.Context,
	_ *empty.Empty,
) (*vagrant_plugin_sdk.Project_TargetIdsResponse, error) {
	ids, err := p.Impl.TargetIds()
	if err != nil {
		return nil, err
	}

	return &vagrant_plugin_sdk.Project_TargetIdsResponse{
		Ids: ids}, nil
}

var (
	_ plugin.Plugin     = (*ProjectPlugin)(nil)
	_ plugin.GRPCPlugin = (*ProjectPlugin)(nil)
	_ core.Project      = (*projectClient)(nil)
)