package core

import "io"

type NamedPlugin struct {
	Plugin  interface{}
	Name    string
	Type    string
	Options interface{}
}

type PluginManager interface {
	ListPlugins(types ...string) (plugins []*NamedPlugin, err error)
	GetPlugin(name, typeName string) (*NamedPlugin, error)

	io.Closer
}
