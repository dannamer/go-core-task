package main

import (
	"testing"
	"reflect"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

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

	mergedChannel := mergeChannels([]chan int{ch1, ch2})

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	var result []int
	for val := range mergedChannel {
		result = append(result, val)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
