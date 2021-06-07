package core

// import (
// 	"github.com/hashicorp/vagrant-plugin-sdk/docs"
// )
import "github.com/hashicorp/go-argmapper"

type Host interface {
	// Config() interface{}
	// Documentation() (*docs.Documentation, error)
	Detect() bool
	HasCapability(name string) bool
	Capability(name string, args ...argmapper.Arg) (interface{}, error)
}
