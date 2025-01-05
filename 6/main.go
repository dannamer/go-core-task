package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateRandomNumbers(ch chan int, n int, wg *sync.WaitGroup) {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	defer wg.Done()

	for i := 0; i < n; i++ {
		ch <- rng.Intn(100)
	}
	close(ch)
}

func readNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go generateRandomNumbers(ch, 5, &wg)

	wg.Add(1)
	go readNumbers(ch, &wg)

	wg.Wait()
}
