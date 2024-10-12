package chapterone

import (
	"fmt"
	"sync"
)

func MemoryAccessSynchronization1() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Println("the values is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
}

func MemoryAccessSynchronization2() {
	var memoryAccess sync.Mutex
	var value int
	go func() {
		memoryAccess.Lock()
		value++
		memoryAccess.Unlock()
	}()

	memoryAccess.Lock()
	if value == 0 {
		fmt.Printf("the value is %v.\n", value)
	} else {
		fmt.Printf("the value is %v.\n", value)
	}
	memoryAccess.Unlock()
}
