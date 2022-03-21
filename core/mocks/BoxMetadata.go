// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	mock "github.com/stretchr/testify/mock"
)

// BoxMetadata is an autogenerated mock type for the BoxMetadata type
type BoxMetadata struct {
	mock.Mock
}

// ListProviders provides a mock function with given fields: version
func (_m *BoxMetadata) ListProviders(version string) ([]string, error) {
	ret := _m.Called(version)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(version)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(version)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListVersions provides a mock function with given fields: opts
func (_m *BoxMetadata) ListVersions(opts ...*core.BoxProvider) ([]string, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []string
	if rf, ok := ret.Get(0).(func(...*core.BoxProvider) []string); ok {
		r0 = rf(opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...*core.BoxProvider) error); ok {
		r1 = rf(opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Matches provides a mock function with given fields: version, name, provider
func (_m *BoxMetadata) Matches(version string, name string, provider *core.BoxProvider) (bool, error) {
	ret := _m.Called(version, name, provider)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, *core.BoxProvider) bool); ok {
		r0 = rf(version, name, provider)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *core.BoxProvider) error); ok {
		r1 = rf(version, name, provider)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MatchesAny provides a mock function with given fields: version, name, provider
func (_m *BoxMetadata) MatchesAny(version string, name string, provider ...*core.BoxProvider) (bool, error) {
	_va := make([]interface{}, len(provider))
	for _i := range provider {
		_va[_i] = provider[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, version, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, ...*core.BoxProvider) bool); ok {
		r0 = rf(version, name, provider...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, ...*core.BoxProvider) error); ok {
		r1 = rf(version, name, provider...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *BoxMetadata) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Provider provides a mock function with given fields: version, name
func (_m *BoxMetadata) Provider(version string, name string) (*core.BoxProvider, error) {
	ret := _m.Called(version, name)

	var r0 *core.BoxProvider
	if rf, ok := ret.Get(0).(func(string, string) *core.BoxProvider); ok {
		r0 = rf(version, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BoxProvider)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(version, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Version provides a mock function with given fields: version, opts
func (_m *BoxMetadata) Version(version string, opts *core.BoxProvider) (*core.BoxVersion, error) {
	ret := _m.Called(version, opts)

	var r0 *core.BoxVersion
	if rf, ok := ret.Get(0).(func(string, *core.BoxProvider) *core.BoxVersion); ok {
		r0 = rf(version, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BoxVersion)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *core.BoxProvider) error); ok {
		r1 = rf(version, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
