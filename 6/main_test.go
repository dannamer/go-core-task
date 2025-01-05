package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestGenerateRandomNumbers(t *testing.T) {
	tests := []struct {
		n           int
		expectedMin int
		expectedMax int
	}{
		{n: 5, expectedMin: 0, expectedMax: 99},
		{n: 10, expectedMin: 0, expectedMax: 99},
		{n: 0, expectedMin: 0, expectedMax: 0},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Generating %d numbers", test.n), func(t *testing.T) {
			ch := make(chan int)
			var wg sync.WaitGroup
			wg.Add(1)

			go generateRandomNumbers(ch, test.n, &wg)

			count := 0
			for num := range ch {
				if num < test.expectedMin || num > test.expectedMax {
					t.Errorf("generated number %d is out of expected range [%d, %d]", num, test.expectedMin, test.expectedMax)
				}
				count++
			}
			if count != test.n {
				t.Errorf("expected %d numbers, got %d", test.n, count)
			}
			select {
			case _, ok := <-ch:
				if ok {
					t.Errorf("channel should be closed after generation")
				}
			default:
			}
		})
	}
}
