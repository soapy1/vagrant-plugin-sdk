package core

import (
	"io"
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	"github.com/hashicorp/vagrant-plugin-sdk/datadir"
	"github.com/hashicorp/vagrant-plugin-sdk/helper/path"
	"github.com/hashicorp/vagrant-plugin-sdk/terminal"
)

type Target interface {
	Name() (string, error)
	SetName(value string) (err error)
	ResourceId() (string, error)
	Project() (Project, error)
	Metadata() (map[string]string, error)
	DataDir() (*datadir.Target, error)
	State() (State, error)
	UI() (ui terminal.UI, err error)
	VagrantfileName() (name string, err error)
	VagrantfilePath() (p path.Path, err error)
	UpdatedAt() (t *time.Time, err error)

	Provider() (p Provider, err error)
	Communicate() (comm Communicator, err error)

	Record() (*anypb.Any, error)
	Specialize(kind interface{}) (specialized Machine, err error) // TODO(spox): mapping needs to be fixed so return is interface{}

	io.Closer
}