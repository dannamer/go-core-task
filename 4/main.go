package main

import (
	"fmt"
)

func difference(slice1, slice2 []string) []string {
	lookup := make(map[string]struct{})
	for _, item := range slice2 {
		lookup[item] = struct{}{}
	}

	var result []string
	for _, item := range slice1 {
		if _, found := lookup[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

func main() {
	slice1 := []string{"apple", "apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	result := difference(slice1, slice2)

	fmt.Println(result)
}
