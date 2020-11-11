// Package component has the interfaces for all the components that
// can be implemented. A component is the broad term used to describe
// all providers, provisioners, etc.
//
// Many component interfaces have functions named `XFunc` where "X" is some
// operation and the return value is "interface{}". These functions should return
// a method handle to the function implementing that operation. This pattern is
// done so that we can support custom typed operations that take and return
// full rich types for an operation. We use a minimal dependency-injection
// framework (see internal/mapper) to call these functions.
package component

//go:generate stringer -type=Type -linecomment
//go:generate mockery -all -case underscore

// Type is an enum of all the types of components supported.
// This isn't used directly in this package but is used by other packages
// to reference the component types.
type Type uint

const (
	InvalidType     Type = iota // Invalid
	ProviderType                // Provider
	ProvisionerType             // Provisioner
	maxType
)

// TypeMap is a mapping of Type to the nil pointer to the interface of that
// type. This can be used with libraries such as mapper.
var TypeMap = map[Type]interface{}{
	ProviderType:    (*Provider)(nil),
	ProvisionerType: (*Provisioner)(nil),
}

// Providers are the backend that VMs are launched on
type Provider interface {
	// Handles operations involving interfacing with a provider
	ProviderFunc() interface{}
}

// Provisioner is responsible for provisioning a VM
type Provisioner interface {
	// Handles operations involving provisioining the guest machine
	ProvisionerFunc() interface{}
}

type LabelSet struct {
	Labels map[string]string
}

// JobInfo is available to plugins to get information about the context
// in which a job is executing.
type JobInfo struct {
	// Id is the ID of the job that is executing this plugin operation.
	// If this is empty then it means that the execution is happening
	// outside of a job.
	Id string

	// Local is true if the operation is running locally on a machine
	// alongside the invocation. This can be used to determine if you can
	// do things such as open browser windows, read user files, etc.
	Local bool

	// Workspace is the workspace that this job is executing in. This should
	// be used by plugins to properly isolate resources from each other.
	Workspace string
}

// See Args.Source in the protobuf protocol.
type Source struct {
	App  string
	Path string
}

