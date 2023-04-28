// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xslices

import (
	"github.com/manchurio/xgo/internal/constraints"
	"github.com/manchurio/xgo/internal/irand"
	"sort"
)

func ToMap[K comparable, T any](slice []T, f func(T) K) map[K]T {
	m := make(map[K]T, len(slice))
	for i := range slice {
		v := slice[i]
		k := f(v)
		m[k] = v
	}
	return m
}

func Shuffle[T any](s []T) {
	rnd := irand.GetRand()
	rnd.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	irand.PutRand(rnd)
}

func Sort[T any](s []T, f func(i, j int) bool) {
	sort.Slice(s, f)
}

func Sum[T constraints.Number](s []T) (sum T) {
	for _, v := range s {
		sum += v
	}
	return sum
}

// Last Get the last value of the slice, return zero-value when the length is 0
func Last[T any](s []T) T {
	l, _ := LastB(s)
	return l
}

// LastB Get the last value of the slice, return zero-value,false when the length is 0
func LastB[T any](s []T) (T, bool) {
	if l := len(s); l > 0 {
		return s[l-1], true
	}
	var t T
	return t, false
}

// First Get the first value of the slice, return zero-value when the length is 0
func First[T any](s []T) T {
	l, _ := FirstB(s)
	return l
}

// FirstB Get the first value of the slice, return zero-value,false when the length is 0
func FirstB[T any](s []T) (T, bool) {
	if l := len(s); l > 0 {
		return s[0], true
	}
	var t T
	return t, false
}
