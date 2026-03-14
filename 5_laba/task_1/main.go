package main

import (
	"fmt"
	"sync"
)

var COUNTER int32
var mtx sync.Mutex

func GenerateNumber(evenChanel chan int, oddChanel chan int) {
	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenChanel <- i
		} else {
			oddChanel <- i
		}
	}
	close(evenChanel)
	close(oddChanel)
}

func main() {
	evenChannel := make(chan int)
	oddChannel := make(chan int)

	go GenerateNumber(evenChannel, oddChannel)

	for evenChannel != nil || oddChannel != nil {
		select {
		case number, ok := <-evenChannel:
			if !ok {
				evenChannel = nil
				continue
			}
			if number%3 == 0 {
				mtx.Lock()
				COUNTER++
				mtx.Unlock()
			}

		case number, ok := <-oddChannel:
			if !ok {
				oddChannel = nil
				continue
			}
			if number%33 == 0 {
				mtx.Lock()
				COUNTER--
				mtx.Unlock()
			}
		}
	}

	fmt.Println(COUNTER)
}
