package main

import (
	"concurrency-in-go/chapterone"
	"concurrency-in-go/chapterthree"
	"os"
)

func main() {
	sample := os.Getenv("SAMPLE_FIRST_PAGE")
	switch sample {
	case "4":
		chapterone.RaceCondition()
	case "5":
		chapterone.RaceCondition2()
	case "8-1":
		chapterone.MemoryAccessSynchronization1()
	case "8-2":
		chapterone.MemoryAccessSynchronization2()
	case "9":
		chapterone.Deadlock()
	case "14":
		chapterone.Livelock()
	case "16":
		chapterone.Starvation()
	case "40":
		chapterthree.JoinPoint()
	case "41":
		chapterthree.JoinPointWithArray()
	case "43":
		chapterthree.CountGoRoutinesSize()
	case "47":
		chapterthree.WaitGroup()
	case "49":
		chapterthree.Mutext()
	case "51":
		chapterthree.RWMutex()
	}
}
