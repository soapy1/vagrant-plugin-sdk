// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	anypb "google.golang.org/protobuf/types/known/anypb"

	datadir "github.com/hashicorp/vagrant-plugin-sdk/datadir"

	mock "github.com/stretchr/testify/mock"

	terminal "github.com/hashicorp/vagrant-plugin-sdk/terminal"

	time "time"
)

// Machine is an autogenerated mock type for the Machine type
type Machine struct {
	mock.Mock
}

// Box provides a mock function with given fields:
func (_m *Machine) Box() (core.Box, error) {
	ret := _m.Called()

	var r0 core.Box
	if rf, ok := ret.Get(0).(func() core.Box); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Box)
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

// Close provides a mock function with given fields:
func (_m *Machine) Close() error {
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
func (_m *Machine) Communicate() (core.Communicator, error) {
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

// ConnectionInfo provides a mock function with given fields:
func (_m *Machine) ConnectionInfo() (*core.ConnectionInfo, error) {
	ret := _m.Called()

	var r0 *core.ConnectionInfo
	if rf, ok := ret.Get(0).(func() *core.ConnectionInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.ConnectionInfo)
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
func (_m *Machine) DataDir() (*datadir.Target, error) {
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
func (_m *Machine) Destroy() error {
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
func (_m *Machine) GetUUID() (string, error) {
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

// Guest provides a mock function with given fields:
func (_m *Machine) Guest() (core.Guest, error) {
	ret := _m.Called()

	var r0 core.Guest
	if rf, ok := ret.Get(0).(func() core.Guest); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Guest)
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

// ID provides a mock function with given fields:
func (_m *Machine) ID() (string, error) {
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

// Inspect provides a mock function with given fields:
func (_m *Machine) Inspect() (string, error) {
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

// MachineState provides a mock function with given fields:
func (_m *Machine) MachineState() (*core.MachineState, error) {
	ret := _m.Called()

	var r0 *core.MachineState
	if rf, ok := ret.Get(0).(func() *core.MachineState); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.MachineState)
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

// Metadata provides a mock function with given fields:
func (_m *Machine) Metadata() (map[string]string, error) {
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
func (_m *Machine) Name() (string, error) {
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
func (_m *Machine) Project() (core.Project, error) {
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
func (_m *Machine) Provider() (core.Provider, error) {
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
func (_m *Machine) ProviderName() (string, error) {
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
func (_m *Machine) Record() (*anypb.Any, error) {
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
func (_m *Machine) ResourceId() (string, error) {
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
func (_m *Machine) Save() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetID provides a mock function with given fields: value
func (_m *Machine) SetID(value string) error {
	ret := _m.Called(value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetMachineState provides a mock function with given fields: state
func (_m *Machine) SetMachineState(state *core.MachineState) error {
	ret := _m.Called(state)

	var r0 error
	if rf, ok := ret.Get(0).(func(*core.MachineState) error); ok {
		r0 = rf(state)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetName provides a mock function with given fields: value
func (_m *Machine) SetName(value string) error {
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
func (_m *Machine) SetUUID(id string) error {
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
func (_m *Machine) Specialize(kind interface{}) (interface{}, error) {
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
func (_m *Machine) State() (core.State, error) {
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

// SyncedFolders provides a mock function with given fields:
func (_m *Machine) SyncedFolders() ([]*core.MachineSyncedFolder, error) {
	ret := _m.Called()

	var r0 []*core.MachineSyncedFolder
	if rf, ok := ret.Get(0).(func() []*core.MachineSyncedFolder); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.MachineSyncedFolder)
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

// UI provides a mock function with given fields:
func (_m *Machine) UI() (terminal.UI, error) {
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

// UID provides a mock function with given fields:
func (_m *Machine) UID() (string, error) {
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

// UpdatedAt provides a mock function with given fields:
func (_m *Machine) UpdatedAt() (*time.Time, error) {
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

type NewMachineT interface {
	mock.TestingT
	Cleanup(func())
}

// NewMachine creates a new instance of Machine. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMachine(t NewMachineT) *Machine {
	mock := &Machine{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
