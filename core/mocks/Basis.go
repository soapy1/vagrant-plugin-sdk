// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	datadir "github.com/hashicorp/vagrant-plugin-sdk/datadir"

	mock "github.com/stretchr/testify/mock"

	path "github.com/hashicorp/vagrant-plugin-sdk/helper/path"

	terminal "github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// Basis is an autogenerated mock type for the Basis type
type Basis struct {
	mock.Mock
}

// Boxes provides a mock function with given fields:
func (_m *Basis) Boxes() (core.BoxCollection, error) {
	ret := _m.Called()

	var r0 core.BoxCollection
	if rf, ok := ret.Get(0).(func() core.BoxCollection); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.BoxCollection)
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

// CWD provides a mock function with given fields:
func (_m *Basis) CWD() (path.Path, error) {
	ret := _m.Called()

	var r0 path.Path
	if rf, ok := ret.Get(0).(func() path.Path); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Path)
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
func (_m *Basis) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DataDir provides a mock function with given fields:
func (_m *Basis) DataDir() (*datadir.Basis, error) {
	ret := _m.Called()

	var r0 *datadir.Basis
	if rf, ok := ret.Get(0).(func() *datadir.Basis); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datadir.Basis)
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

// DefaultPrivateKey provides a mock function with given fields:
func (_m *Basis) DefaultPrivateKey() (path.Path, error) {
	ret := _m.Called()

	var r0 path.Path
	if rf, ok := ret.Get(0).(func() path.Path); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(path.Path)
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

// Host provides a mock function with given fields:
func (_m *Basis) Host() (core.Host, error) {
	ret := _m.Called()

	var r0 core.Host
	if rf, ok := ret.Get(0).(func() core.Host); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Host)
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
func (_m *Basis) ResourceId() (string, error) {
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

// TargetIndex provides a mock function with given fields:
func (_m *Basis) TargetIndex() (core.TargetIndex, error) {
	ret := _m.Called()

	var r0 core.TargetIndex
	if rf, ok := ret.Get(0).(func() core.TargetIndex); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.TargetIndex)
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
func (_m *Basis) UI() (terminal.UI, error) {
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

type NewBasisT interface {
	mock.TestingT
	Cleanup(func())
}

// NewBasis creates a new instance of Basis. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBasis(t NewBasisT) *Basis {
	mock := &Basis{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
