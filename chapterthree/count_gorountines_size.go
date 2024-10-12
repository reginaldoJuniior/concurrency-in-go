package chapterthree

import (
	"fmt"
	"runtime"
	"sync"
)

func CountGoRoutinesSize() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // We require a goroutine that will never exit so that we can keep a number of them.

	const numGoroutines = 1e4 // Here we define the number of goroutines to create. We will use the law of large numbers to asymptotically approach the size of a goroutine.
	wg.Add(numGoroutines)
	before := memConsumed() // Here we measure the amount of memory consumed before creating the goroutines.
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()
	after := memConsumed() // And here we measure the amount of memory consumed after creating the goroutines.
	fmt.Printf("%.3fkb\n", float64(after-before)/numGoroutines/1000)
}
