// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xmath

import (
	"github.com/manchurio/xgo/internal/constraints"
	"math"
)

// Floor returns the greatest integer value less than or equal to x.
//
// Special cases are:
//
//	Floor(±0) = ±0
//	Floor(±Inf) = ±Inf
//	Floor(NaN) = NaN
func Floor[F constraints.Float, I constraints.Integer](f F) I {
	return I(math.Floor(float64(f)))
}

func FloorInt[F constraints.Float](f F) int {
	return int(math.Floor(float64(f)))
}

func FloorInt32[F constraints.Float](f F) int32 {
	return int32(math.Floor(float64(f)))
}

func FloorInt64[F constraints.Float](f F) int32 {
	return int32(math.Floor(float64(f)))
}

func Ceil[F constraints.Float, I constraints.Integer](f F) I {
	return I(math.Ceil(float64(f)))
}

func CeilInt[F constraints.Float](f F) int {
	return Ceil[F, int](f)
}

func CeilInt64[F constraints.Float](f F) int64 {
	return Ceil[F, int64](f)
}

func CeilInt32[F constraints.Float](f F) int32 {
	return Ceil[F, int32](f)
}
