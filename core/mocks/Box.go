// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	mock "github.com/stretchr/testify/mock"
)

// Box is an autogenerated mock type for the Box type
type Box struct {
	mock.Mock
}

// AutomaticUpdateCheckAllowed provides a mock function with given fields:
func (_m *Box) AutomaticUpdateCheckAllowed() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Compare provides a mock function with given fields: box
func (_m *Box) Compare(box core.Box) (int, error) {
	ret := _m.Called(box)

	var r0 int
	if rf, ok := ret.Get(0).(func(core.Box) int); ok {
		r0 = rf(box)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.Box) error); ok {
		r1 = rf(box)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Destroy provides a mock function with given fields:
func (_m *Box) Destroy() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Directory provides a mock function with given fields:
func (_m *Box) Directory() (string, error) {
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

// HasUpdate provides a mock function with given fields: version
func (_m *Box) HasUpdate(version string) (bool, error) {
	ret := _m.Called(version)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(version)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InUse provides a mock function with given fields: index
func (_m *Box) InUse(index core.TargetIndex) (bool, error) {
	ret := _m.Called(index)

	var r0 bool
	if rf, ok := ret.Get(0).(func(core.TargetIndex) bool); ok {
		r0 = rf(index)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.TargetIndex) error); ok {
		r1 = rf(index)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Metadata provides a mock function with given fields:
func (_m *Box) Metadata() (core.BoxMetadataMap, error) {
	ret := _m.Called()

	var r0 core.BoxMetadataMap
	if rf, ok := ret.Get(0).(func() core.BoxMetadataMap); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.BoxMetadataMap)
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

// MetadataURL provides a mock function with given fields:
func (_m *Box) MetadataURL() (string, error) {
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

// Name provides a mock function with given fields:
func (_m *Box) Name() (string, error) {
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

// Provider provides a mock function with given fields:
func (_m *Box) Provider() (string, error) {
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

// Repackage provides a mock function with given fields: path
func (_m *Box) Repackage(path string) error {
	ret := _m.Called(path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Version provides a mock function with given fields:
func (_m *Box) Version() (string, error) {
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
