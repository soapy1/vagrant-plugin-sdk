// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Push is an autogenerated mock type for the Push type
type Push struct {
	mock.Mock
}

// Push provides a mock function with given fields:
func (_m *Push) Push() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}