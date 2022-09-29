package adapters

import (
	"math"

	"github.com/dosovma/otus_arch/pkg"
)

type MovableAdapter struct {
	Obj UObject
}

func (m MovableAdapter) GetPosition() (pkg.Vector, error) {
	p, ok := m.Obj.GetProperty("position")
	if !ok {
		return nil, ErrNotImplemented
	}

	position, ok := p.(pkg.Vector)
	if !ok {
		return nil, ErrInvalidProperty
	}

	return position, nil
}

func (m MovableAdapter) GetVelocity() (pkg.Vector, error) {
	d, ok := m.Obj.GetProperty("direction")
	if !ok {
		return nil, ErrNotImplemented
	}

	direction, ok := d.(int)
	if !ok {
		return nil, ErrInvalidProperty
	}

	md, ok := m.Obj.GetProperty("maxDirections")
	if !ok {
		return nil, ErrNotImplemented
	}

	maxDirections, ok := md.(int)
	if !ok {
		return nil, ErrInvalidProperty
	}

	v, ok := m.Obj.GetProperty("velocity")
	if !ok {
		return nil, ErrNotImplemented
	}

	velocity, ok := v.(int)
	if !ok {
		return nil, ErrInvalidProperty
	}

	return pkg.Vector{
		velocity * int(math.Cos(2*math.Pi*float64(direction)/float64(maxDirections))),
		velocity * int(math.Sin(2*math.Pi*float64(direction)/float64(maxDirections))),
	}, nil
}

func (m MovableAdapter) SetPosition(position pkg.Vector) error {
	if ok := m.Obj.SetProperty("position", position); !ok {
		return ErrNotImplemented
	}

	return nil
}
