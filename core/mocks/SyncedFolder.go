// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	mock "github.com/stretchr/testify/mock"
)

// SyncedFolder is an autogenerated mock type for the SyncedFolder type
type SyncedFolder struct {
	mock.Mock
}

// Capability provides a mock function with given fields: name, args
func (_m *SyncedFolder) Capability(name string, args ...interface{}) (interface{}, error) {
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, args...)
	ret := _m.Called(_ca...)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string, ...interface{}) interface{}); ok {
		r0 = rf(name, args...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...interface{}) error); ok {
		r1 = rf(name, args...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Cleanup provides a mock function with given fields: machine, opts
func (_m *SyncedFolder) Cleanup(machine core.Machine, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, machine)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, ...interface{}) error); ok {
		r0 = rf(machine, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Disable provides a mock function with given fields: machine, folders, opts
func (_m *SyncedFolder) Disable(machine core.Machine, folders []*core.Folder, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, machine, folders)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, []*core.Folder, ...interface{}) error); ok {
		r0 = rf(machine, folders, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Enable provides a mock function with given fields: machine, folders, opts
func (_m *SyncedFolder) Enable(machine core.Machine, folders []*core.Folder, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, machine, folders)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, []*core.Folder, ...interface{}) error); ok {
		r0 = rf(machine, folders, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HasCapability provides a mock function with given fields: name
func (_m *SyncedFolder) HasCapability(name string) (bool, error) {
	ret := _m.Called(name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Prepare provides a mock function with given fields: machine, folders, opts
func (_m *SyncedFolder) Prepare(machine core.Machine, folders []*core.Folder, opts ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, machine, folders)
	_ca = append(_ca, opts...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, []*core.Folder, ...interface{}) error); ok {
		r0 = rf(machine, folders, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Seed provides a mock function with given fields: _a0
func (_m *SyncedFolder) Seed(_a0 *core.Seeds) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*core.Seeds) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Seeds provides a mock function with given fields:
func (_m *SyncedFolder) Seeds() (*core.Seeds, error) {
	ret := _m.Called()

	var r0 *core.Seeds
	if rf, ok := ret.Get(0).(func() *core.Seeds); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Seeds)
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

// Usable provides a mock function with given fields: machine
func (_m *SyncedFolder) Usable(machine core.Machine) (bool, error) {
	ret := _m.Called(machine)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.Machine) bool); ok {
		r0 = rf(machine)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.Machine) error); ok {
		r1 = rf(machine)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewSyncedFolderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewSyncedFolder creates a new instance of SyncedFolder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewSyncedFolder(t NewSyncedFolderT) *SyncedFolder {
	mock := &SyncedFolder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
