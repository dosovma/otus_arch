package app_test

import (
	"math"
	"reflect"
	"testing"

	"github.com/dosovma/otus_arch/internal/app"
)

func Test_Solve(t *testing.T) {
	type args struct {
		a float64
		b float64
		c float64
	}
	tests := []struct {
		name      string
		args      args
		wantRoots []float64
		wantErr   bool
	}{
		{
			name: "case: x^2 + 1 = 0: there are no roots",
			args: args{
				a: 1,
				b: 0,
				c: 1,
			},
			wantRoots: []float64{},
			wantErr:   false,
		},
		{
			name: "case: x^2 - 1 = 0: there are two different roots",
			args: args{
				a: 1,
				b: 0,
				c: -1,
			},
			wantRoots: []float64{1, -1},
			wantErr:   false,
		},
		{
			name: "case: x^2 + 2x + 1 = 0: there is only one root",
			args: args{
				a: 1,
				b: 2,
				c: 1,
			},
			wantRoots: []float64{-1},
			wantErr:   false,
		},
		{
			name: "case: discriminant is very small: there is only one root",
			args: args{
				a: 1.999999999999999,
				b: 4,
				c: 1.999999999999999,
			},
			wantRoots: []float64{-3.999999999999998},
			wantErr:   false,
		},
		{
			name: "case: a = 0: invalid quadratic equation",
			args: args{
				a: 0,
				b: 2,
				c: 1,
			},
			wantRoots: nil,
			wantErr:   true,
		},
		{
			name: "case: a is very small: invalid quadratic equation",
			args: args{
				a: 0.0000000000000001,
				b: 2,
				c: 1,
			},
			wantRoots: nil,
			wantErr:   true,
		},
		{
			name: "case: a is not number: invalid quadratic equation",
			args: args{
				a: math.NaN(),
				b: 2,
				c: 4,
			},
			wantRoots: nil,
			wantErr:   true,
		},
		{
			name: "case: b is positive infinity: invalid quadratic equation",
			args: args{
				a: 1,
				b: math.Inf(1), // positive infinity
				c: 0,
			},
			wantRoots: nil,
			wantErr:   true,
		},
		{
			name: "case: c is negative infinity: invalid quadratic equation",
			args: args{
				a: 1,
				b: 4,
				c: math.Inf(-1), // negative infinity
			},
			wantRoots: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRoots, gotErr := app.Solve(tt.args.a, tt.args.b, tt.args.c)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("solve() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRoots, tt.wantRoots) {
				t.Errorf("solve() got = %v, wantX1 %v", gotRoots, tt.wantRoots)
			}
		})
	}
}
