// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	component "github.com/hashicorp/vagrant-plugin-sdk/component"
	mock "github.com/stretchr/testify/mock"
)

// PluginInfo is an autogenerated mock type for the PluginInfo type
type PluginInfo struct {
	mock.Mock
}

// ComponentTypes provides a mock function with given fields:
func (_m *PluginInfo) ComponentTypes() []component.Type {
	ret := _m.Called()

	var r0 []component.Type
	if rf, ok := ret.Get(0).(func() []component.Type); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]component.Type)
		}
	}

	return r0
}