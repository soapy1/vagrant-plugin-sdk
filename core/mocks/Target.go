// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	anypb "google.golang.org/protobuf/types/known/anypb"

	datadir "github.com/hashicorp/vagrant-plugin-sdk/datadir"

	mock "github.com/stretchr/testify/mock"

	terminal "github.com/hashicorp/vagrant-plugin-sdk/terminal"

	time "time"
)

// Target is an autogenerated mock type for the Target type
type Target struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *Target) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Communicate provides a mock function with given fields:
func (_m *Target) Communicate() (core.Communicator, error) {
	ret := _m.Called()

	var r0 core.Communicator
	if rf, ok := ret.Get(0).(func() core.Communicator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Communicator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DataDir provides a mock function with given fields:
func (_m *Target) DataDir() (*datadir.Target, error) {
	ret := _m.Called()

	var r0 *datadir.Target
	if rf, ok := ret.Get(0).(func() *datadir.Target); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datadir.Target)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields:
func (_m *Target) Destroy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetUUID provides a mock function with given fields:
func (_m *Target) GetUUID() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Metadata provides a mock function with given fields:
func (_m *Target) Metadata() (map[string]string, error) {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *Target) Name() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Project provides a mock function with given fields:
func (_m *Target) Project() (core.Project, error) {
	ret := _m.Called()

	var r0 core.Project
	if rf, ok := ret.Get(0).(func() core.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Project)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Provider provides a mock function with given fields:
func (_m *Target) Provider() (core.Provider, error) {
	ret := _m.Called()

	var r0 core.Provider
	if rf, ok := ret.Get(0).(func() core.Provider); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Provider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ProviderName provides a mock function with given fields:
func (_m *Target) ProviderName() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Record provides a mock function with given fields:
func (_m *Target) Record() (*anypb.Any, error) {
	ret := _m.Called()

	var r0 *anypb.Any
	if rf, ok := ret.Get(0).(func() *anypb.Any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*anypb.Any)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResourceId provides a mock function with given fields:
func (_m *Target) ResourceId() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields:
func (_m *Target) Save() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetName provides a mock function with given fields: value
func (_m *Target) SetName(value string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetUUID provides a mock function with given fields: id
func (_m *Target) SetUUID(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Specialize provides a mock function with given fields: kind
func (_m *Target) Specialize(kind interface{}) (interface{}, error) {
	ret := _m.Called(kind)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(interface{}) interface{}); ok {
		r0 = rf(kind)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}) error); ok {
		r1 = rf(kind)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// State provides a mock function with given fields:
func (_m *Target) State() (core.State, error) {
	ret := _m.Called()

	var r0 core.State
	if rf, ok := ret.Get(0).(func() core.State); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(core.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UI provides a mock function with given fields:
func (_m *Target) UI() (terminal.UI, error) {
	ret := _m.Called()

	var r0 terminal.UI
	if rf, ok := ret.Get(0).(func() terminal.UI); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(terminal.UI)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatedAt provides a mock function with given fields:
func (_m *Target) UpdatedAt() (*time.Time, error) {
	ret := _m.Called()

	var r0 *time.Time
	if rf, ok := ret.Get(0).(func() *time.Time); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*time.Time)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Vagrantfile provides a mock function with given fields:
func (_m *Target) Vagrantfile() (core.Vagrantfile, error) {
	ret := _m.Called()

	var r0 core.Vagrantfile
	if rf, ok := ret.Get(0).(func() core.Vagrantfile); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Vagrantfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewTargetT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTarget creates a new instance of Target. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTarget(t NewTargetT) *Target {
	mock := &Target{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
