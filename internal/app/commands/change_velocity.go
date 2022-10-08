package commands

//go:generate mockgen -destination=./mocks/change_velocity.go -package=mocks -source=change_velocity.go
type IChangeVelocity interface {
	IsMovable() (bool, error)
	ChangeVelocity() error
}

type ChangeVelocityCmd struct {
	ChangeVelocity IChangeVelocity
}

// Execute of ChangeVelocityCmd changes velocity. If an object doesn't move, execute doesn't change the velocity.
func (cv ChangeVelocityCmd) Execute() error {
	isMovable, err := cv.ChangeVelocity.IsMovable()
	if err != nil {
		return ErrInvalidOperation
	}

	if !isMovable {
		return nil
	}

	return cv.ChangeVelocity.ChangeVelocity()
}
