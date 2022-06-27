// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Configurable is an autogenerated mock type for the Configurable type
type Configurable struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *Configurable) Config() (interface{}, error) {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
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

type NewConfigurableT interface {
	mock.TestingT
	Cleanup(func())
}

// NewConfigurable creates a new instance of Configurable. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewConfigurable(t NewConfigurableT) *Configurable {
	mock := &Configurable{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
