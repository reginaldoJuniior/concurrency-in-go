package chapterthree

import (
	"fmt"
	"sync"
	"time"
)

func WaitGroup() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		println("It is goroutine sleeping...")
		time.Sleep(1 * time.Second)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		println("It is the 2nd goroutine sleeping...")
		time.Sleep(2 * time.Second)
	}()

	wg.Wait()
	fmt.Println("All goroutines complete.")
}
