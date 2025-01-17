package core

type TargetIndex interface {
	All() (targets []Target, err error)
	Delete(uuid string) (err error)
	Get(uuid string) (entry Target, err error)
	Includes(uuid string) (exists bool, err error)
	Set(entry Target) (updatedEntry Target, err error)
	// Recover(entry Target) (updatedEntry Target, err error)
}
