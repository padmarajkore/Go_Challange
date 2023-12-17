package main

import "fmt"

func Facotrial(n int) int {
	f := 1

	if n == 0 {
		return 1
	}
	
	for x :=1; x <= n; x++ {
		f = f*x
	}
	return f
}

func main (){
	fmt.Println(Facotrial(3))
	fmt.Println(Facotrial(0))
}