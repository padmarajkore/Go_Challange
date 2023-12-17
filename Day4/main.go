package main

import "fmt"

func IsLeap(year int) bool {
	
	if year%4 == 0{
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

func main() {

	fmt.Println(IsLeap(2024))
	fmt.Println(IsLeap(2000))
	fmt.Println(IsLeap(1900))
	fmt.Println(IsLeap(2022))

}