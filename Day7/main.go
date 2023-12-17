package main

import "fmt"

func PositiveSum(slice []int) int {

	positivesum:= 0
	for i:=0; i < len(slice); i++ {
		if slice[i] > 0 {
			positivesum += slice[i]
		}
	}
	return positivesum
}

func main() {
    numbers := []int{3, -5, 7, 0, 1, -2, 8}
    fmt.Println(PositiveSum(numbers)) // Should print 19
}