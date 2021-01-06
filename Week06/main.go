package main

import (
	"fmt"
	"sync"
	"time"
)

var currentTime = time.Now()

func Counter() {
	counter := &rollingCounter{
		mu: sync.Mutex{},
		window: newWindow(&windowOpt{
			size: 10,
		}),
		size:     10,
		duration: time.Second,
		offset:   0,
		nowFunc:  TimeFunc,
		lastIncr: currentTime,
	}
	for i := 0; i < 9; i++ {
		counter.Incr(Success)
	}
	if count := counter.GetCurrentCount(Success); count != 9 {
		//t.Errorf("Error: actual:%d expected: %d", count, 9)
	}
	for i := 0; i < 5; i++ {
		counter.Incr(Success)
	}
	if count := counter.GetCurrentCount(Success); count != 5 {
		//t.Errorf("Error: actual:%d expected: %d", count, 5)
	}
}

func TimeFunc() time.Time {
	currentTime = currentTime.Add(time.Millisecond * 100)
	return currentTime
}

func main() {
	fmt.Println("go week06")
	Counter()

	select {}
}
