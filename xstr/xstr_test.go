// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xstr

import (
	"github.com/manchurio/xgo/internal/constraints"
	"testing"
)

func TestToInt(t *testing.T) {
	type args[T constraints.Integer] struct {
		s      string
		defVal T
	}
	type testCase[T constraints.Integer] struct {
		name string
		args args[T]
		want T
	}
	{
		tests := []testCase[int]{
			{args: args[int]{s: "1", defVal: 10}, want: 1},
			{args: args[int]{s: "1.2", defVal: 1}, want: 1},
			{args: args[int]{s: "ab", defVal: 1}, want: 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ToInt(tt.args.s, tt.args.defVal); got != tt.want {
					t.Errorf("ToInt() = %v, want %v", got, tt.want)
				}
			})
		}
	}
	{
		tests := []testCase[int8]{
			{args: args[int8]{s: "-1", defVal: 10}, want: -1},
			{args: args[int8]{s: "1.2", defVal: 1}, want: 1},
			{args: args[int8]{s: "ab", defVal: 1}, want: 1},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := ToInt(tt.args.s, tt.args.defVal); got != tt.want {
					t.Errorf("ToInt() = %v, want %v", got, tt.want)
				}
			})
		}
	}
}

func TestMustToInt(t *testing.T) {
	type args struct {
		s string
	}
	type testCase struct {
		name string
		args args
		want int
	}
	tests := []testCase{
		{args: args{s: "1.0"}, want: 1},
		{args: args{s: "2"}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {

				}
			}()
			if got := MustToInt[int](tt.args.s); got != tt.want {
				t.Errorf("MustToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
