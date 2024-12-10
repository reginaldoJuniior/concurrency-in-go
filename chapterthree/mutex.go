package chapterthree

import "sync"

func Mutex() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()         // here we request exclusive use of the critical section - in this case, the count variable
		defer lock.Unlock() // here we indicate that we are done with the critical section
		count++
		println("Incrementing:", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		println("Decrementing:", count)
	}

	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}

	arithmetic.Wait()
	println("Arithmetic complete.")
}
