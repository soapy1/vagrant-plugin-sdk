// Code generated by mockery 2.12.3. DO NOT EDIT.

package mocks

import (
	component "github.com/hashicorp/vagrant-plugin-sdk/component"
	core "github.com/hashicorp/vagrant-plugin-sdk/core"

	mock "github.com/stretchr/testify/mock"
)

// Provisioner is an autogenerated mock type for the Provisioner type
type Provisioner struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields: machine, config
func (_m *Provisioner) Cleanup(machine core.Machine, config *component.ConfigData) error {
	ret := _m.Called(machine, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, *component.ConfigData) error); ok {
		r0 = rf(machine, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Configure provides a mock function with given fields: machine, config, rootConfig
func (_m *Provisioner) Configure(machine core.Machine, config *component.ConfigData, rootConfig *component.ConfigData) error {
	ret := _m.Called(machine, config, rootConfig)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, *component.ConfigData, *component.ConfigData) error); ok {
		r0 = rf(machine, config, rootConfig)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Provision provides a mock function with given fields: machine, config
func (_m *Provisioner) Provision(machine core.Machine, config *component.ConfigData) error {
	ret := _m.Called(machine, config)

	var r0 error
	if rf, ok := ret.Get(0).(func(core.Machine, *component.ConfigData) error); ok {
		r0 = rf(machine, config)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type NewProvisionerT interface {
	mock.TestingT
	Cleanup(func())
}

// NewProvisioner creates a new instance of Provisioner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProvisioner(t NewProvisionerT) *Provisioner {
	mock := &Provisioner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
