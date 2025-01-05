package main

import "fmt"

func findIntersection(a, b []int) (bool, []int) {
	lookup := make(map[int]struct{})
	for _, num := range a {
		lookup[num] = struct{}{}
	}

	var intersection []int
	for _, num := range b {
		if _, found := lookup[num]; found {
			intersection = append(intersection, num)
		}
	}

	return len(intersection) > 0, intersection
}

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	intersectionExist, intersection := findIntersection(a, b)
	fmt.Println(intersectionExist)
	fmt.Println(intersection)
}