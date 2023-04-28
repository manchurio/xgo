// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xslices

import (
	"github.com/manchurio/xgo/internal/constraints"
	"github.com/manchurio/xgo/irand"
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
