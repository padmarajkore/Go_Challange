package main

import "fmt"

func UniqueElements(s []int) []int {

	uniqueMap := make(map[int]bool)

	uniqueSlice := []int{}

	for _, num := range s {
		if !uniqueMap[num] {
			uniqueMap[num] = true
			uniqueSlice = append(uniqueSlice, num)
		}
	}
	return uniqueSlice
}

func main() {

	fmt.Println(UniqueElements([]int{3, 4, 5, 6, 4, 5, 2, 3, 4}))

}
