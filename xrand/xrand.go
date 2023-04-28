// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xrand

import (
	"github.com/manchurio/xgo/constraints"
	"math/rand"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() any {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func getRand() *rand.Rand {
	return pool.Get().(*rand.Rand)
}

// RandInt Returns a random int within [startInclusive,endInclusive]
func RandInt[T constraints.Integer](startInclusive, endInclusive T) T {
	r := getRand()
	res := r.Int63n(int64(endInclusive-startInclusive+1)) + int64(startInclusive)
	t := T(res)
	pool.Put(r)
	return t
}

// RandIntN Returns a random int within [startInclusive,endExclusive)
func RandIntN[T constraints.Integer](startInclusive, endExclusive T) T {
	return RandInt(startInclusive, endExclusive-1)
}
