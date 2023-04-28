// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xstr

import (
	"github.com/manchurio/xgo/constraints"
	"strconv"
)

func ToInt[T constraints.Integer](s string, defVal T) T {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		return defVal
	}
	return T(atoi)
}

func MustToInt[T constraints.Integer](s string) T {
	atoi, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return T(atoi)
}
