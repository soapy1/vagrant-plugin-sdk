package core

import (
	"github.com/hashicorp/vagrant-plugin-sdk/component"
)

type Vagrantfile interface {
	GetRootConfig() (*component.ConfigData, error)
	GetConfig(namespace string) (*component.ConfigData, error)
	GetValue(path ...string) (interface{}, error)
	PrimaryTargetName() (name string, err error)
	SetConfig(namespace string, config *component.ConfigData) error
	Target(name, provider string) (Target, error)
	TargetConfig(name, provider string, validateProvider bool) (Vagrantfile, error)
	TargetNames() (names []string, err error)
	//TargetNamesAndOptions() (names []string, options map[string]interface{}, err error)
}
