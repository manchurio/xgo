// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xmath

import (
	"github.com/manchurio/xgo/internal/constraints"
	"reflect"
	"testing"
)

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		x T
		y []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	{
		tests := []testCase[int]{
			{args: args[int]{x: 1, y: []int{2, 3}}, want: 1},
			{args: args[int]{x: -1, y: []int{0, 1}}, want: -1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Min(tt.args.x, tt.args.y...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	{
		tests := []testCase[float64]{
			{args: args[float64]{x: 1, y: []float64{2, 3}}, want: 1},
			{args: args[float64]{x: -1, y: []float64{0, 1}}, want: -1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Min(tt.args.x, tt.args.y...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestMax(t *testing.T) {
	type args[T constraints.Ordered] struct {
		x T
		y []T
	}
	type testCase[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	{
		tests := []testCase[int]{
			{args: args[int]{x: 1, y: []int{2, 3}}, want: 3},
			{args: args[int]{x: -1, y: []int{0, 1}}, want: 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Max(tt.args.x, tt.args.y...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	{
		tests := []testCase[float64]{
			{args: args[float64]{x: 1, y: []float64{2, 3}}, want: 3},
			{args: args[float64]{x: -1, y: []float64{0, 1}}, want: 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := Max(tt.args.x, tt.args.y...); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("Min() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestCompare(t *testing.T) {
	type args[T constraints.Number] struct {
		x T
		y T
	}
	type testCase[T constraints.Number] struct {
		name string
		args args[T]
		want int
	}
	tests := []testCase[int]{
		{args: args[int]{1, 2}, want: -1},
		{args: args[int]{2, 1}, want: 1},
		{args: args[int]{2, 2}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Compare(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}
