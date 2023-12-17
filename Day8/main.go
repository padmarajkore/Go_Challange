package main

import "fmt"

func ProductNonNegative(slice []int) int {

	if len(slice) == 0{
		return 1
	}

	product:= 1
	for i :=0; i < len(slice); i++ {
		if slice[i] > 0 {
			product = product * slice[i]
		}
	}

	return product
}

func main() {
	fmt.Println(ProductNonNegative([]int{4,5,6,-6,7}))
}