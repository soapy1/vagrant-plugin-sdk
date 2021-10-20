package core

import (
	"io"
)

type Guest interface {
	// Config() interface{}
	// Documentation() (*docs.Documentation, error)
	Capability(name string, args ...interface{}) (interface{}, error)
	Detect(Target) (bool, error)
	HasCapability(name string) (bool, error)
	Parent() (string, error)

	io.Closer
}
