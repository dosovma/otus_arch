package adapters

type ChangeVelocityAdapter struct {
	Obj UObject
}

func (cv ChangeVelocityAdapter) IsMovable() (bool, error) {
	return false, nil
}

func (cv ChangeVelocityAdapter) ChangeVelocity() error {
	return nil
}
