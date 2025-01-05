package main

import (
	"fmt"
	"sync"
)

func mergeChannels(channels []chan int) chan int {
	merged := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			for val := range c {
				merged <- val
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()
	go func() {
		for i := 10; i < 15; i++ {
			ch3 <- i
		}
		close(ch3)
	}()

	mergedChannel := mergeChannels([]chan int{ch1, ch2, ch3})
	for val := range mergedChannel {
		fmt.Println(val)
	}
}
