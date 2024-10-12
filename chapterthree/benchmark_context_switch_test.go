package chapterthree

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // Here we wait until we're told to begin.
		for i := 0; i < b.N; i++ {
			c <- token // Here we send messages to the receiver.
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // Here we wait until we're told to begin.
		for i := 0; i < b.N; i++ {
			<-c // Here we receive messages from the sender.
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // Here we begin the performance timer.
	close(begin)   // Here we tell the two goroutines to begin.
	wg.Wait()
}
