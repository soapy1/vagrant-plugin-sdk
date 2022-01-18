// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	core "github.com/hashicorp/vagrant-plugin-sdk/core"
	datadir "github.com/hashicorp/vagrant-plugin-sdk/datadir"

	mock "github.com/stretchr/testify/mock"

	path "github.com/hashicorp/vagrant-plugin-sdk/helper/path"

	terminal "github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

// Project is an autogenerated mock type for the Project type
type Project struct {
	mock.Mock
}

// Boxes provides a mock function with given fields:
func (_m *Project) Boxes() (core.BoxCollection, error) {
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
func (_m *Project) CWD() (string, error) {
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

// Close provides a mock function with given fields:
func (_m *Project) Close() error {
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
func (_m *Project) DataDir() (*datadir.Project, error) {
	ret := _m.Called()

	var r0 *datadir.Project
	if rf, ok := ret.Get(0).(func() *datadir.Project); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datadir.Project)
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
func (_m *Project) DefaultPrivateKey() (string, error) {
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

// Home provides a mock function with given fields:
func (_m *Project) Home() (string, error) {
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

// Host provides a mock function with given fields:
func (_m *Project) Host() (core.Host, error) {
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

// LocalData provides a mock function with given fields:
func (_m *Project) LocalData() (string, error) {
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

// ResourceId provides a mock function with given fields:
func (_m *Project) ResourceId() (string, error) {
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

// Target provides a mock function with given fields: name
func (_m *Project) Target(name string) (core.Target, error) {
	ret := _m.Called(name)

	var r0 core.Target
	if rf, ok := ret.Get(0).(func(string) core.Target); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(core.Target)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TargetIds provides a mock function with given fields:
func (_m *Project) TargetIds() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// TargetIndex provides a mock function with given fields:
func (_m *Project) TargetIndex() (core.TargetIndex, error) {
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

// TargetNames provides a mock function with given fields:
func (_m *Project) TargetNames() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
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

// Tmp provides a mock function with given fields:
func (_m *Project) Tmp() (string, error) {
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

// UI provides a mock function with given fields:
func (_m *Project) UI() (terminal.UI, error) {
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

// VagrantfileName provides a mock function with given fields:
func (_m *Project) VagrantfileName() (string, error) {
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

// VagrantfilePath provides a mock function with given fields:
func (_m *Project) VagrantfilePath() (path.Path, error) {
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
