package adapters

import "github.com/dosovma/otus_arch/internal/app/entity"

type ChangeVelocityAdapter struct {
	Obj entity.UObject
}

func (cv ChangeVelocityAdapter) IsMovable() (bool, error) {
	return false, nil
}

func (cv ChangeVelocityAdapter) ChangeVelocity() error {
	return nil
}
