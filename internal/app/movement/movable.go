package movement

import (
	"github.com/dosovma/otus_arch/pkg"
)

//go:generate mockgen -destination=./mocks/movable.go -package=mocks -source=movable.go
type Movable interface {
	GetPosition() (pkg.Vector, error)
	GetVelocity() (pkg.Vector, error)
	SetPosition(position pkg.Vector) error
}

type Move struct {
	Move Movable
}

func (m Move) Execute() error {
	currentPosition, err := m.Move.GetPosition()
	if err != nil {
		return ErrInvalidOperation
	}

	velocity, err := m.Move.GetVelocity()
	if err != nil {
		return ErrInvalidOperation
	}

	if len(velocity) != len(currentPosition) {
		return ErrInvalidOperation
	}

	newPosition := make(pkg.Vector, len(currentPosition))
	for i := range currentPosition {
		newPosition[i] = currentPosition[i] + velocity[i]
	}

	return m.Move.SetPosition(newPosition)
}
