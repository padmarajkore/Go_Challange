package main

import "fmt"

func FindMax(s []int) int {

	if len(s) == 0 {
		return -1
	}

	MaxIndex := 0

	for i:=1; i < len(s); i++ {
		if s[i] > s[MaxIndex] {
			MaxIndex = i
		}
	}
	return s[MaxIndex]
}

func main (){
	fmt.Println(FindMax([]int{3,4,5,6,9,5,6,7}))
}