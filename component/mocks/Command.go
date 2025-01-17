// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Command is an autogenerated mock type for the Command type
type Command struct {
	mock.Mock
}

// CommandInfoFunc provides a mock function with given fields:
func (_m *Command) CommandInfoFunc() interface{} {
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

// ExecuteFunc provides a mock function with given fields: _a0
func (_m *Command) ExecuteFunc(_a0 []string) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func([]string) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

type NewCommandT interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommand creates a new instance of Command. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommand(t NewCommandT) *Command {
	mock := &Command{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
