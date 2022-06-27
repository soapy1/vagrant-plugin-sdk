// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// CapabilityPlatform is an autogenerated mock type for the CapabilityPlatform type
type CapabilityPlatform struct {
	mock.Mock
}

// CapabilityFunc provides a mock function with given fields: capName
func (_m *CapabilityPlatform) CapabilityFunc(capName string) interface{} {
	ret := _m.Called(capName)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(string) interface{}); ok {
		r0 = rf(capName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// HasCapabilityFunc provides a mock function with given fields:
func (_m *CapabilityPlatform) HasCapabilityFunc() interface{} {
	ret := _m.Called()

	var r0 interface{}
	if rf, ok := ret.Get(0).(func() interface{}); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type NewCapabilityPlatformT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCapabilityPlatform creates a new instance of CapabilityPlatform. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCapabilityPlatform(t NewCapabilityPlatformT) *CapabilityPlatform {
	mock := &CapabilityPlatform{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
