// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xmath

import (
	"github.com/manchurio/xgo/internal/constraints"
	"testing"
)

func TestFloor(t *testing.T) {
	type args[F constraints.Float] struct {
		f F
	}
	type testCase[F constraints.Float, I constraints.Integer] struct {
		name string
		args args[F]
		want I
	}
	tests := []testCase[float64, int32]{
		{args: args[float64]{f: 1.32}, want: 1},
		{args: args[float64]{f: 2.32}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Floor[float64, int32](tt.args.f); got != tt.want {
				t.Errorf("Floor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCeilInt64(t *testing.T) {
	type args[F constraints.Float] struct {
		f F
	}
	type testCase[F constraints.Float] struct {
		name string
		args args[F]
		want int64
	}
	tests := []testCase[float64]{
		{args: args[float64]{f: 3.14}, want: 4},
		{args: args[float64]{f: 0.14}, want: 1},
		{args: args[float64]{f: 0.00}, want: 0},
		{args: args[float64]{f: -10.1}, want: -10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CeilInt64(tt.args.f); got != tt.want {
				t.Errorf("CeilInt64() = %v, want %v", got, tt.want)
			}
		})
	}
}
