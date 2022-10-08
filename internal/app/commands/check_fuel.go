package commands

import "errors"

var ErrEmptyFuel = errors.New("there is no enough fuel")

//go:generate mockgen -destination=./mocks/check_fuel.go -package=mocks -source=check_fuel.go
type ICheckFuel interface {
	GetFuel() (int, error)
}

type CheckFuelCmd struct {
	CheckFuel ICheckFuel
}

func (cf CheckFuelCmd) Execute() error {
	fuel, err := cf.CheckFuel.GetFuel()
	if err != nil {
		return ErrInvalidOperation
	}

	if fuel <= 0 {
		return ErrEmptyFuel
	}

	return nil
}
