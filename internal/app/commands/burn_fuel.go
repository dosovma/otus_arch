package commands

//go:generate mockgen -destination=./mocks/burn_fuel.go -package=mocks -source=burn_fuel.go
type IBurnFuel interface {
	GetFuel() (int, error)
	GetFuelConsumption() (int, error)
	SetFuel(fuel int) error
}

type BurnFuelCmd struct {
	BurnFuel IBurnFuel
}

func (bf BurnFuelCmd) Execute() error {
	fuel, err := bf.BurnFuel.GetFuel()
	if err != nil {
		return ErrInvalidOperation
	}

	fuelConsumption, err := bf.BurnFuel.GetFuelConsumption()
	if err != nil {
		return ErrInvalidOperation
	}

	return bf.BurnFuel.SetFuel(fuel - fuelConsumption)
}
