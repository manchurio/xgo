// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package irand

import (
	"math/rand"
	"sync"
	"time"
)

var pool = sync.Pool{
	New: func() any {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

func GetRand() *rand.Rand {
	return pool.Get().(*rand.Rand)
}

func PutRand(x *rand.Rand) {
	pool.Put(x)
}
