package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func generateRandomSlice() []int {
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	originalSlice := make([]int, 10)
	for i := 0; i < len(originalSlice); i++ {
		originalSlice[i] = randGen.Intn(100)
	}
	return originalSlice
}

func sliceExample(s []int) []int {
	var evenSlice []int
	for _, v := range s {
		if v%2 == 0 {
			evenSlice = append(evenSlice, v)
		}
	}
	return evenSlice
}

func addElements(s []int, num int) []int {
	return append(s, num)
}

func copySlice(s []int) []int {
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	return newSlice
}

func removeElement(s []int, index int) ([]int, error) {
	if index < 0 || index >= len(s) {
		return nil, errors.New("index out of range")
	}
	return append(s[:index], s[index+1:]...), nil
}

func main() {
	originalSlice := generateRandomSlice()
	fmt.Println("Original Slice:", originalSlice)

	evenSlice := sliceExample(originalSlice)
	fmt.Println("Even Numbers:", evenSlice)

	newSlice := addElements(originalSlice, 42)
	fmt.Println("After Adding 42:", newSlice)

	copiedSlice := copySlice(originalSlice)
	fmt.Println("Copied Slice:", copiedSlice)

	removedSlice, _ := removeElement(originalSlice, 2)
	fmt.Println("After Removing Index 2:", removedSlice)
}
