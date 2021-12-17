// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	mock "github.com/stretchr/testify/mock"
)

// PluginManager is an autogenerated mock type for the PluginManager type
type PluginManager struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *PluginManager) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPlugin provides a mock function with given fields: name, typeName
func (_m *PluginManager) GetPlugin(name string, typeName string) (*core.NamedPlugin, error) {
	ret := _m.Called(name, typeName)

	var r0 *core.NamedPlugin
	if rf, ok := ret.Get(0).(func(string, string) *core.NamedPlugin); ok {
		r0 = rf(name, typeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.NamedPlugin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(name, typeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPlugins provides a mock function with given fields: types
func (_m *PluginManager) ListPlugins(types ...string) ([]*core.NamedPlugin, error) {
	_va := make([]interface{}, len(types))
	for _i := range types {
		_va[_i] = types[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*core.NamedPlugin
	if rf, ok := ret.Get(0).(func(...string) []*core.NamedPlugin); ok {
		r0 = rf(types...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*core.NamedPlugin)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(types...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
