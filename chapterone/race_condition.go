package chapterone

import (
	"fmt"
	"time"
)

func RaceCondition() {
	var data int
	go func() {
		data++
	}()
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}

func RaceCondition2() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1 * time.Second) // this is bad
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
