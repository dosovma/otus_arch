package commands

//go:generate mockgen -destination=./mocks/rotatable.go -package=mocks -source=rotatable.go
type Rotatable interface {
	GetDirection() (int, error)
	GetAngularVelocity() (int, error)
	SetDirection(direction int) error
	GetMaxDirections() (int, error)
}

type Rotate struct {
	Rotate Rotatable
}

func (rm Rotate) Execute() error {
	currentDirection, err := rm.Rotate.GetDirection()
	if err != nil {
		return ErrInvalidOperation
	}

	angVelocity, err := rm.Rotate.GetAngularVelocity()
	if err != nil {
		return ErrInvalidOperation
	}

	maxDirections, err := rm.Rotate.GetMaxDirections()
	if err != nil {
		return ErrInvalidOperation
	}

	return rm.Rotate.SetDirection((currentDirection + angVelocity) % maxDirections)
}
