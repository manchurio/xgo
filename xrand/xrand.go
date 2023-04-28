// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xrand

import (
	"github.com/manchurio/xgo/internal/constraints"
	"github.com/manchurio/xgo/internal/irand"
)

// RandInt Returns a random int within [startInclusive,endInclusive]
func RandInt[T constraints.Integer](startInclusive, endInclusive T) T {
	r := irand.GetRand()
	res := r.Int63n(int64(endInclusive-startInclusive+1)) + int64(startInclusive)
	t := T(res)
	irand.PutRand(r)
	return t
}

// RandIntN Returns a random int within [startInclusive,endExclusive)
func RandIntN[T constraints.Integer](startInclusive, endExclusive T) T {
	return RandInt(startInclusive, endExclusive-1)
}
