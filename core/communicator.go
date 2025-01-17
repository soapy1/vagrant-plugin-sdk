package core

type CommunicatorMessage struct {
	ExitCode int32
	Stdout   string
	Stderr   string
}

type Communicator interface {
	Seeder

	// Config() interface{}
	// Documentation() (*docs.Documentation, error)
	Download(machine Machine, source, destination string) error
	Execute(machine Machine, command []string, opts ...interface{}) (status int32, err error)
	Init(machine Machine) error
	Match(machine Machine) (isMatch bool, err error)
	PrivilegedExecute(machine Machine, command []string, opts ...interface{}) (status int32, err error)
	Ready(machine Machine) (isReady bool, err error)
	Reset(machine Machine) error
	Test(machine Machine, command []string, opts ...interface{}) (valid bool, err error)
	Upload(machine Machine, source, destination string) error
	WaitForReady(machine Machine, wait int) (isReady bool, err error)
}
