package core

import (
	"context"
	"errors"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/hashicorp/go-argmapper"
	"github.com/hashicorp/go-plugin"

	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/datadir"
	vplugin "github.com/hashicorp/vagrant-plugin-sdk/internal/plugin"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

type TargetPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl core.Target
	*vplugin.BasePlugin
}

// Implements plugin.GRPCPlugin
func (p *TargetPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &targetClient{
		client:     vagrant_plugin_sdk.NewTargetServiceClient(c),
		BaseClient: p.NewClient(ctx, broker, nil),
	}, nil
}

func (p *TargetPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	vagrant_plugin_sdk.RegisterTargetServiceServer(s, &targetServer{
		Impl:       p.Impl,
		BaseServer: p.NewServer(broker, nil),
	})
	return nil
}

// Target implements core.Target interface
type targetClient struct {
	*vplugin.BaseClient

	client vagrant_plugin_sdk.TargetServiceClient
}

type targetServer struct {
	*vplugin.BaseServer

	Impl core.Target
	vagrant_plugin_sdk.UnimplementedTargetServiceServer
}

func (c *targetClient) Communicate() (comm core.Communicator, err error) {
	commArg, err := c.client.Communicate(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to call communicator from client, %w", err)
	}
	result, err := c.Map(
		commArg,
		(*core.Communicator)(nil),
		argmapper.Typed(c.Ctx),
		argmapper.Typed(c.Logger),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to map communicator, %w", err)
	}
	comm = result.(core.Communicator)
	return
}

func (c *targetClient) SetName(name string) (err error) {
	_, err = c.client.SetName(c.Ctx, &vagrant_plugin_sdk.Target_SetNameRequest{
		Name: name})
	return
}

func (c *targetClient) Provider() (p core.Provider, err error) {
	pr, err := c.client.Provider(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	result, err := c.Map(pr, (*core.Provider)(nil), argmapper.Typed(c.Ctx))
	if err != nil {
		return
	}
	p = result.(core.Provider)
	return
}

func (c *targetClient) ProviderName() (name string, err error) {
	result, err := c.client.ProviderName(c.Ctx, &emptypb.Empty{})
	if err == nil {
		name = result.Name
	}

	return
}

func (c *targetClient) UpdatedAt() (utime *time.Time, err error) {
	r, err := c.client.UpdatedAt(c.Ctx, &emptypb.Empty{})
	if err == nil {
		ut := r.UpdatedAt.AsTime()
		utime = &ut
	}

	return
}

func (c *targetClient) Name() (name string, err error) {
	r, err := c.client.Name(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	name = r.Name
	return
}

func (c *targetClient) ResourceId() (rid string, err error) {
	r, err := c.client.ResourceId(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	rid = r.ResourceId

	return
}

func (c *targetClient) Project() (project core.Project, err error) {
	r, err := c.client.Project(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	result, err := c.Map(r, (*core.Project)(nil), argmapper.Typed(c.Ctx))
	if err != nil {
		return
	}
	project = result.(core.Project)
	return
}

func (c *targetClient) Metadata() (mdata map[string]string, err error) {
	r, err := c.client.Metadata(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	result, err := c.Map(r, (*map[string]string)(nil))
	if err != nil {
		return
	}
	mdata = result.(map[string]string)
	return
}

func (c *targetClient) DataDir() (dir *datadir.Target, err error) {
	r, err := c.client.DataDir(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	result, err := c.Map(r, (*datadir.Target)(nil))
	if err != nil {
		return
	}
	dir = result.(*datadir.Target)
	return
}

func (c *targetClient) State() (state core.State, err error) {
	r, err := c.client.State(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	result, err := c.Map(r, (*core.State)(nil))
	if err != nil {
		return
	}
	state = result.(core.State)
	return
}

func (c *targetClient) Record() (record *anypb.Any, err error) {
	r, err := c.client.Record(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	record = r.Record
	return
}

func (c *targetClient) GetUUID() (id string, err error) {
	uuid, err := c.client.GetUUID(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}
	id = uuid.Uuid
	return
}

func (c *targetClient) SetUUID(uuid string) (err error) {
	_, err = c.client.SetUUID(
		c.Ctx,
		&vagrant_plugin_sdk.Target_SetUUIDRequest{
			Uuid: uuid,
		},
	)
	return
}

func (c *targetClient) Specialize(kind interface{}) (specialized interface{}, err error) {
	// TODO: specialize type based on the `kind` requested
	a, err := anypb.New(&emptypb.Empty{})
	if err != nil {
		return
	}
	r, err := c.client.Specialize(c.Ctx, a)

	if err != nil {
		return
	}

	m := &vagrant_plugin_sdk.Args_Target_Machine{}
	if err = r.UnmarshalTo(m); err != nil {
		return
	}

	s, err := c.Map(m, (*core.Machine)(nil),
		argmapper.Typed(c.Ctx))
	return s.(core.Machine), err
}

func (c *targetClient) UI() (ui terminal.UI, err error) {
	r, err := c.client.UI(c.Ctx, &emptypb.Empty{})
	if err != nil {
		return
	}

	result, err := c.Map(r, (*terminal.UI)(nil),
		argmapper.Typed(c.Ctx))
	if err == nil {
		ui = result.(terminal.UI)
	}

	return
}

func (t *targetClient) Save() (err error) {
	_, err = t.client.Save(t.Ctx, &emptypb.Empty{})
	return
}

func (t *targetClient) Destroy() (err error) {
	_, err = t.client.Destroy(t.Ctx, &emptypb.Empty{})
	return
}

func (t *targetClient) Vagrantfile() (core.Vagrantfile, error) {
	resp, err := t.client.Vagrantfile(t.Ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	raw, err := t.Map(resp, (*core.Vagrantfile)(nil), argmapper.Typed(t.Ctx))
	if err != nil {
		return nil, err
	}

	return raw.(core.Vagrantfile), nil
}

// Target Server

func (s *targetServer) Communicate(
	ctx context.Context,
	_ *emptypb.Empty,
) (_ *vagrant_plugin_sdk.Args_Communicator, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target communicator",
				"error", err,
			)
		}
	}()

	c, err := s.Impl.Communicate()
	if err != nil {
		return nil, fmt.Errorf("error getting the communicator, %w", err)
	}

	result, err := s.Map(c, (**vagrant_plugin_sdk.Args_Communicator)(nil), argmapper.Typed(ctx))
	if err != nil {
		return nil, fmt.Errorf("failed to map communicator to proto, %w", err)
	}

	return result.(*vagrant_plugin_sdk.Args_Communicator), nil
}

func (s *targetServer) Name(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Target_NameResponse, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target name",
				"error", err,
			)
		}
	}()

	n, err := s.Impl.Name()
	if err == nil {
		r = &vagrant_plugin_sdk.Target_NameResponse{Name: n}
	}

	return
}

func (t *targetServer) SetName(
	ctx context.Context,
	in *vagrant_plugin_sdk.Target_SetNameRequest,
) (*emptypb.Empty, error) {
	if err := t.Impl.SetName(in.Name); err != nil {
		t.Logger.Error("failed to set target name",
			"error", err,
		)

		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (t *targetServer) Provider(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Args_Provider, err error) {
	defer func() {
		if err != nil {
			t.Logger.Error("failed to get target provider",
				"error", err,
			)
		}
	}()

	p, err := t.Impl.Provider()
	if err != nil {
		return
	}

	result, err := t.Map(p, (**vagrant_plugin_sdk.Args_Provider)(nil),
		argmapper.Typed(ctx))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_Provider)
	}

	return
}

func (t *targetServer) ProviderName(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Target_NameResponse, err error) {
	defer func() {
		if err != nil {
			t.Logger.Error("failed to get target provider name",
				"error", err,
			)
		}
	}()

	pn, err := t.Impl.ProviderName()
	if err == nil {
		r = &vagrant_plugin_sdk.Target_NameResponse{Name: pn}
	}

	return
}

func (t *targetServer) UpdatedAt(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Target_UpdatedAtResponse, error) {
	u, err := t.Impl.UpdatedAt()
	if err != nil {
		t.Logger.Error("failed to get target updated at time",
			"error", err,
		)

		return nil, err
	}

	return &vagrant_plugin_sdk.Target_UpdatedAtResponse{
		UpdatedAt: timestamppb.New(*u)}, nil
}

func (s *targetServer) ResourceId(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Target_ResourceIdResponse, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target resource id",
				"error", err,
			)
		}
	}()

	rid, err := s.Impl.ResourceId()
	if err == nil {
		r = &vagrant_plugin_sdk.Target_ResourceIdResponse{ResourceId: rid}
	}

	return
}

func (s *targetServer) Project(
	ctx context.Context,
	_ *emptypb.Empty,
) (_ *vagrant_plugin_sdk.Args_Project, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target project",
				"error", err,
			)
		}
	}()

	p, err := s.Impl.Project()
	if err != nil {
		return nil, err
	}

	result, err := s.Map(p, (**vagrant_plugin_sdk.Args_Project)(nil))
	if err != nil {
		return nil, err
	}
	return result.(*vagrant_plugin_sdk.Args_Project), nil
}

func (s *targetServer) Metadata(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Args_MetadataSet, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target metadata",
				"error", err,
			)
		}
	}()

	m, err := s.Impl.Metadata()
	if err != nil {
		return
	}
	result, err := s.Map(m, (**vagrant_plugin_sdk.Args_MetadataSet)(nil))
	if err != nil {
		return
	}
	r = result.(*vagrant_plugin_sdk.Args_MetadataSet)

	return
}

func (s *targetServer) DataDir(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Args_DataDir_Target, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target data directory",
				"error", err,
			)
		}
	}()

	d, err := s.Impl.DataDir()
	if err != nil {
		return
	}
	if d != nil {
		var result interface{}
		result, err = s.Map(d, (**vagrant_plugin_sdk.Args_DataDir_Target)(nil))
		if err != nil {
			return nil, err
		}
		r = result.(*vagrant_plugin_sdk.Args_DataDir_Target)
	} else {
		r = &vagrant_plugin_sdk.Args_DataDir_Target{}
	}
	return
}

func (s *targetServer) State(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Args_Target_State, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target state",
				"error", err,
			)
		}
	}()

	st, err := s.Impl.State()
	if err != nil {
		return
	}
	result, err := s.Map(st, (**vagrant_plugin_sdk.Args_Target_State)(nil))
	if err != nil {
		return
	}
	r = result.(*vagrant_plugin_sdk.Args_Target_State)

	return
}

func (s *targetServer) Record(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Target_RecordResponse, err error) {
	defer func() {
		if err != nil {
			s.Logger.Error("failed to get target ui",
				"error", err,
			)
		}
	}()

	record, err := s.Impl.Record()
	if err == nil {
		r = &vagrant_plugin_sdk.Target_RecordResponse{Record: record}
	}

	return
}

func (t *targetServer) UI(
	ctx context.Context,
	_ *emptypb.Empty,
) (r *vagrant_plugin_sdk.Args_TerminalUI, err error) {
	defer func() {
		if err != nil {
			t.Logger.Error("failed to get target ui",
				"error", err,
			)
		}
	}()

	d, err := t.Impl.UI()
	if err != nil {
		return
	}

	result, err := t.Map(d, (**vagrant_plugin_sdk.Args_TerminalUI)(nil),
		argmapper.Typed(ctx))
	if err == nil {
		r = result.(*vagrant_plugin_sdk.Args_TerminalUI)
	}

	return
}

func (t *targetServer) SetUUID(
	ctx context.Context,
	in *vagrant_plugin_sdk.Target_SetUUIDRequest,
) (*emptypb.Empty, error) {
	err := t.Impl.SetUUID(in.Uuid)
	if err != nil {
		t.Logger.Error("failed to set target uuid",
			"error", err,
		)
	}
	return &emptypb.Empty{}, err
}

func (t *targetServer) GetUUID(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Target_GetUUIDResponse, error) {
	uuid, err := t.Impl.GetUUID()
	if err != nil {
		t.Logger.Error("failed to get target uuid",
			"error", err,
		)

		return nil, err
	}

	return &vagrant_plugin_sdk.Target_GetUUIDResponse{
		Uuid: uuid}, nil
}

func (t *targetServer) Specialize(
	ctx context.Context,
	in *anypb.Any,
) (r *anypb.Any, err error) {
	defer func() {
		if err != nil {
			t.Logger.Error("failed to specialize target",
				"error", err,
			)
		}
	}()

	mc, ok := t.Impl.(interface{ Machine() core.Machine })
	if !ok {
		err = errors.New("could not specialize to machine")
		return
	}

	result, err := t.Map(mc.Machine(), (**vagrant_plugin_sdk.Args_Target_Machine)(nil),
		argmapper.Typed(ctx))
	if err != nil {
		return
	}
	return anypb.New(result.(*vagrant_plugin_sdk.Args_Target_Machine))
}

func (t *targetServer) Save(
	ctx context.Context,
	_ *emptypb.Empty,
) (_ *emptypb.Empty, err error) {
	err = t.Impl.Save()
	return
}

func (t *targetServer) Destroy(
	ctx context.Context,
	_ *emptypb.Empty,
) (_ *emptypb.Empty, err error) {
	err = t.Impl.Destroy()
	return
}

func (t *targetServer) Vagrantfile(
	ctx context.Context,
	_ *emptypb.Empty,
) (*vagrant_plugin_sdk.Args_Vagrantfile, error) {
	v, err := t.Impl.Vagrantfile()
	if err != nil {
		t.Logger.Error("failed to get vagrantfile from target implementation",
			"error", err,
		)
		return nil, err
	}

	raw, err := t.Map(v, (**vagrant_plugin_sdk.Args_Vagrantfile)(nil), argmapper.Typed(ctx))
	if err != nil {
		return nil, err
	}

	return raw.(*vagrant_plugin_sdk.Args_Vagrantfile), nil
}

var (
	_ plugin.Plugin                          = (*TargetPlugin)(nil)
	_ plugin.GRPCPlugin                      = (*TargetPlugin)(nil)
	_ core.Target                            = (*targetClient)(nil)
	_ vagrant_plugin_sdk.TargetServiceServer = (*targetServer)(nil)
)
