package commands

// BaseCmd needs to start process and testing.
type BaseCmd struct{}

// Execute throws an error.
func (b BaseCmd) Execute() error {
	return ErrConnectionTimeout
}
