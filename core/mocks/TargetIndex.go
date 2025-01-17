// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	mock "github.com/stretchr/testify/mock"
)

// TargetIndex is an autogenerated mock type for the TargetIndex type
type TargetIndex struct {
	mock.Mock
}

// All provides a mock function with given fields:
func (_m *TargetIndex) All() ([]core.Target, error) {
	ret := _m.Called()

	var r0 []core.Target
	if rf, ok := ret.Get(0).(func() []core.Target); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]core.Target)
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

// Delete provides a mock function with given fields: uuid
func (_m *TargetIndex) Delete(uuid string) error {
	ret := _m.Called(uuid)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: uuid
func (_m *TargetIndex) Get(uuid string) (core.Target, error) {
	ret := _m.Called(uuid)

	var r0 core.Target
	if rf, ok := ret.Get(0).(func(string) core.Target); ok {
		r0 = rf(uuid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Target)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Includes provides a mock function with given fields: uuid
func (_m *TargetIndex) Includes(uuid string) (bool, error) {
	ret := _m.Called(uuid)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(uuid)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(uuid)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: entry
func (_m *TargetIndex) Set(entry core.Target) (core.Target, error) {
	ret := _m.Called(entry)

	var r0 core.Target
	if rf, ok := ret.Get(0).(func(core.Target) core.Target); ok {
		r0 = rf(entry)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Target)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(core.Target) error); ok {
		r1 = rf(entry)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewTargetIndexT interface {
	mock.TestingT
	Cleanup(func())
}

// NewTargetIndex creates a new instance of TargetIndex. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTargetIndex(t NewTargetIndexT) *TargetIndex {
	mock := &TargetIndex{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
