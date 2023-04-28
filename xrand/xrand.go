// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xrand

import (
	"github.com/manchurio/xgo/internal/constraints"
	"math/rand"
)

// RandInt Returns a random int within startInclusive - endExclusive-1
func RandInt[T constraints.Integer](startInclusive, endExclusive T) T {
	res := rand.Int63n(int64(endExclusive-startInclusive+1)) + int64(startInclusive)
	return T(res)
}

// RandIntN Returns a random int within startInclusive - endInclusive
func RandIntN[T constraints.Integer](startInclusive, endInclusive T) T {
	res := rand.Int63n(int64(endInclusive-startInclusive)) + int64(startInclusive)
	return T(res)
}
