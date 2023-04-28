// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xmath

import (
	"github.com/manchurio/xgo/internal/constraints"
)

// Min Returns the minimum value in parameters.
func Min[T constraints.Ordered](x T, y ...T) T {
	var min = x
	for _, v := range y {
		if min > v {
			min = v
		}
	}
	return min
}

// Max Gets the maximum of parameters.
func Max[T constraints.Ordered](x T, y ...T) T {
	var max = x
	for _, v := range y {
		if max < v {
			max = v
		}
	}
	return max
}

// Compare compares two int values numerically
// Returns:  the value 0 if x == y; a value less than 0 if x < y; and a value greater than 0 if x > y
func Compare[T constraints.Number](x, y T) int {
	if x == y {
		return 0
	}
	if x > y {
		return 1
	}
	return -1
}
