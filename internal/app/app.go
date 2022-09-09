package app

import (
	"errors"
	"fmt"
	"math"
)

const inaccuracy = float64(0.0000000001)

// Run is a internal main func to run app.
func Run() {
	var a, b, c float64 = 1, -4, 4

	fmt.Printf("quadratic equation is %fx^2 + %fx + %f\n", a, b, c)

	roots, err := Solve(a, b, c)
	if err != nil {
		fmt.Printf("quadratic equation can not be solved: %s", err.Error())

		return
	}

	switch {
	case len(roots) == 0:
		fmt.Print("there are no roots")
		return
	case len(roots) == 1:
		fmt.Printf("solution is: x1 = %f", roots[0])
		return
	}

	fmt.Printf("solution is: x1 = %f, x2 = %f", roots[0], roots[1])
}

func Solve(a, b, c float64) ([]float64, error) {
	if !isNumberValid(a, b, c) {
		return nil, errors.New("coefficients of equation are not a number")
	}

	if isZero(a) {
		return nil, errors.New("a must be not zero")
	}

	d := b*b - 4*a*c

	if isZero(d) {
		return []float64{-b / 2 * a}, nil
	}

	if isLessZero(d) {
		return []float64{}, nil
	}

	return []float64{
		(-b + math.Sqrt(d)) / 2 * a,
		(-b - math.Sqrt(d)) / 2 * a,
	}, nil
}

func isZero(f float64) bool {
	return f < inaccuracy && f > -inaccuracy
}

func isLessZero(f float64) bool {
	return f < -inaccuracy
}

func isNumberValid(numbers ...float64) bool {
	for _, n := range numbers {
		if math.IsNaN(n) || math.IsInf(n, 1) || math.IsInf(n, -1) {
			return false
		}
	}

	return true
}
