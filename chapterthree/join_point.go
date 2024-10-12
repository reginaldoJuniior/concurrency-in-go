package chapterthree

import (
	"fmt"
	"sync"
)

func JoinPoint() {
	var wg sync.WaitGroup
	salutation := "Hello"
	sayHello := func() {
		wg.Done()
		salutation = "Welcome" // it will replace the value of the variable salutation because it shares the same memory address
		println(salutation)
	}
	wg.Add(1)

	go sayHello()

	wg.Wait() // join point
	fmt.Printf("The value of salutation is %v\n", salutation)
}

func JoinPointWithArray() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}
