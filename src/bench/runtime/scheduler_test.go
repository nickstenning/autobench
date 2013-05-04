// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"testing"
)

func BenchmarkGoroutinesWaiting(b *testing.B) {
	benchmarkGoroutinesWaiting(b, 1000)
}

// benchmarkGoroutinesWaiting opens `simul` goroutines at once, all waiting on a
// value from a single channel.
func benchmarkGoroutinesWaiting(b *testing.B, simul int) {
	c := make(chan bool)
	f := func(c chan bool) { <-c }
	for i := 0; i < b.N; i++ {
		go f(c)
		if (i+1)%simul == 0 {
			c <- true
		}
	}
	c <- true
}
