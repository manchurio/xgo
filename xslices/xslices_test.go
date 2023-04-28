// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xslices_test

import (
	"github.com/manchurio/xgo/internal/constraints"
	"github.com/manchurio/xgo/xslices"
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	type args[T any, K comparable] struct {
		slice []T
		f     func(T) K
	}
	type testCase[T any, K comparable] struct {
		name string
		args args[T, K]
		want map[K]T
	}
	tests := []testCase[float64, int]{
		{args: args[float64, int]{slice: []float64{1, 2}, f: func(f float64) int {
			return int(f) + 1
		}}, want: map[int]float64{2: 1, 3: 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xslices.ToMap(tt.args.slice, tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSort(t *testing.T) {
	var list = []int{9, 7, 8}
	xslices.Sort(list, func(i, j int) bool {
		return list[i] < list[j]
	})
	b := reflect.DeepEqual(list, []int{7, 8, 9})
	if !b {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	type args[T constraints.Number] struct {
		s []T
	}
	type testCase[T constraints.Number] struct {
		name    string
		args    args[T]
		wantSum T
	}
	tests := []testCase[int]{
		{args: args[int]{s: []int{1, 2, 3}}, wantSum: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSum := xslices.Sum(tt.args.s); gotSum != tt.wantSum {
				t.Errorf("Sum() = %v, want %v", gotSum, tt.wantSum)
			}
		})
	}
}

func TestFirstB(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name  string
		args  args[T]
		want  T
		want1 bool
	}
	tests := []testCase[int]{
		{args: args[int]{s: []int{1, 2, 3}}, want: 1, want1: true},
		{args: args[int]{s: []int{}}, want: 0, want1: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := xslices.FirstB(tt.args.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstB() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FirstB() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		s []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{args: args[int]{[]int{1, 2, 3}}, want: 3},
		{args: args[int]{[]int{}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xslices.Last(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}
