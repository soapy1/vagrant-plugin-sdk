package core

import (
	"context"
	"encoding/json"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"

	"github.com/hashicorp/vagrant-plugin-sdk/core"
	"github.com/hashicorp/vagrant-plugin-sdk/proto/vagrant_plugin_sdk"
	"google.golang.org/grpc"
)

type StateBagPlugin struct {
	plugin.NetRPCUnsupportedPlugin

	Impl   core.StateBag
	Logger hclog.Logger
}

func (p *StateBagPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	vagrant_plugin_sdk.RegisterStateBagServiceServer(s, &stateBagServer{
		Impl:   p.Impl,
		Logger: p.Logger,
		Broker: broker,
	})
	return nil
}

func (p *StateBagPlugin) GRPCClient(
	ctx context.Context,
	broker *plugin.GRPCBroker,
	c *grpc.ClientConn,
) (interface{}, error) {
	return &stateBagClient{
		client: vagrant_plugin_sdk.NewStateBagServiceClient(c),
		ctx:    ctx,
		Logger: p.Logger,
		Broker: broker,
	}, nil
}

type stateBagClient struct {
	Logger hclog.Logger
	Broker *plugin.GRPCBroker

	ctx    context.Context
	client vagrant_plugin_sdk.StateBagServiceClient
}

func (c *stateBagClient) Get(key string) (value interface{}) {
	r, err := c.client.Get(c.ctx, &vagrant_plugin_sdk.StateBag_GetRequest{
		Key: key})
	if err != nil {
		c.Logger.Error("failed to get state bag value", "key", key, "error", err)
		return
	}
	err = json.Unmarshal([]byte(r.Value), value)
	if err != nil {
		c.Logger.Error("failed to unmarshal state bag value", "key", key,
			"value", r.Value, "error", err)
	}
	return
}

func (c *stateBagClient) GetOk(key string) (value interface{}, ok bool) {
	r, err := c.client.GetOk(c.ctx, &vagrant_plugin_sdk.StateBag_GetRequest{
		Key: key})
	if err != nil {
		c.Logger.Error("failed to get state bag value", "key", key, "error", err)
		return
	}
	err = json.Unmarshal([]byte(r.Value), value)
	if err != nil {
		c.Logger.Error("failed to unmarshal state bag value", "key", key)
		return
	}
	ok = r.Ok
	return
}

func (c *stateBagClient) Put(key string, value interface{}) {
	v, err := json.Marshal(value)
	_, err = c.client.Put(c.ctx, &vagrant_plugin_sdk.StateBag_PutRequest{
		Key: key, Value: string(v)})
	if err != nil {
		c.Logger.Error("failed to store value in state bag", "key", key,
			"value", value, "error", err)
	}
	return
}

func (c *stateBagClient) Remove(key string) {
	_, err := c.client.Remove(c.ctx, &vagrant_plugin_sdk.StateBag_RemoveRequest{
		Key: key})
	if err != nil {
		c.Logger.Error("failed to remove value from state bag", "key", key,
			"error", err)
	}
	return
}

type stateBagServer struct {
	Logger hclog.Logger
	Broker *plugin.GRPCBroker

	Impl core.StateBag
	vagrant_plugin_sdk.UnimplementedStateBagServiceServer
}

func (s *stateBagServer) Get(
	ctx context.Context,
	req *vagrant_plugin_sdk.StateBag_GetRequest,
) (r *vagrant_plugin_sdk.StateBag_GetResponse, err error) {
	v := s.Impl.Get(req.Key)
	if v == nil {
		v = "null"
	}
	r = &vagrant_plugin_sdk.StateBag_GetResponse{Value: v.(string)}
	return
}

func (s *stateBagServer) GetOk(
	ctx context.Context,
	req *vagrant_plugin_sdk.StateBag_GetRequest,
) (r *vagrant_plugin_sdk.StateBag_GetOkResponse, err error) {
	v, ok := s.Impl.GetOk(req.Key)
	if v == nil {
		v = "null"
	}

	r = &vagrant_plugin_sdk.StateBag_GetOkResponse{
		Ok:    ok,
		Value: v.(string),
	}
	return
}

func (s *stateBagServer) Put(
	ctx context.Context,
	req *vagrant_plugin_sdk.StateBag_PutRequest,
) (r *vagrant_plugin_sdk.StateBag_PutResponse, err error) {
	s.Impl.Put(req.Key, req.Value)
	r = &vagrant_plugin_sdk.StateBag_PutResponse{}
	return
}

func (s *stateBagServer) Remove(
	ctx context.Context,
	req *vagrant_plugin_sdk.StateBag_RemoveRequest,
) (r *vagrant_plugin_sdk.StateBag_RemoveResponse, err error) {
	s.Impl.Remove(req.Key)
	r = &vagrant_plugin_sdk.StateBag_RemoveResponse{}
	return
}

var (
	_ plugin.Plugin                            = (*StateBagPlugin)(nil)
	_ plugin.GRPCPlugin                        = (*StateBagPlugin)(nil)
	_ vagrant_plugin_sdk.StateBagServiceServer = (*stateBagServer)(nil)
	_ core.StateBag                            = (*stateBagClient)(nil)
)
