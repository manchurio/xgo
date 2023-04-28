// Copyright 2023 manchurio. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package xrand

import (
	"sync"
	"testing"
)

func TestRandIntN(t *testing.T) {
	times := 1_000_000
	var start, stop = 10, 20
	var m = make(map[int]int, stop-start)
	for i := 0; i < times; i++ {
		randInt := RandIntN(start, stop)
		m[randInt]++
		if randInt < start || randInt >= stop {
			t.Fatal(randInt)
		}
	}
	t.Logf("m:%v", m)
}

func TestRandInt(t *testing.T) {
	times := 1_000_000
	var start, stop = 10, 20
	var m = make(map[int]int, stop-start)
	for i := 0; i < times; i++ {
		randInt := RandInt(start, stop)
		m[randInt]++
		if randInt < start || randInt > stop {
			t.Fatal(randInt)
		}
	}
	t.Logf("m:%v", m)
}

func TestRandMInt(t *testing.T) {
	times := 100_000
	var start, stop = 10, 20
	var m sync.Map
	var group sync.WaitGroup
	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			for i := 0; i < times; i++ {
				randInt := RandInt(start, stop)
				value, ok := m.Load(randInt)
				if !ok {
					m.Store(randInt, 1)
				} else {
					cnt := value.(int)
					m.Store(randInt, cnt+1)
				}
				if randInt < start || randInt > stop {
					t.Fatal(randInt)
				}
			}
			group.Done()
		}()
	}
	group.Wait()
	var mm = make(map[int]int)
	m.Range(func(key, value any) bool {
		mm[key.(int)] = value.(int)
		return true
	})
	t.Logf("m:%v", mm)
}
