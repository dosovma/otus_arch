package adapters

type RotatableAdapter struct {
	Obj UObject
}

func (r RotatableAdapter) GetDirection() (int, error) {
	p, ok := r.Obj.GetProperty("direction")
	if !ok {
		return 0, ErrNotImplemented
	}

	direction, ok := p.(int)
	if !ok {
		return 0, ErrInvalidProperty
	}

	return direction, nil
}

func (r RotatableAdapter) GetAngularVelocity() (int, error) {
	v, ok := r.Obj.GetProperty("angularVelocity")
	if !ok {
		return 0, ErrNotImplemented
	}

	velocity, ok := v.(int)
	if !ok {
		return 0, ErrInvalidProperty
	}

	return velocity, nil
}

func (r RotatableAdapter) SetDirection(direction int) error {
	if ok := r.Obj.SetProperty("direction", direction); !ok {
		return ErrNotImplemented
	}

	return nil
}

func (r RotatableAdapter) GetMaxDirections() (int, error) {
	md, ok := r.Obj.GetProperty("maxDirections")
	if !ok {
		return 0, ErrNotImplemented
	}

	maxDirections, ok := md.(int)
	if !ok {
		return 0, ErrInvalidProperty
	}

	return maxDirections, nil
}
