package main

import "fmt"

func main() {
	array := [5]int{27, 16, 49, 36, 18}

	var slice []int

	for _, elemento := range array {
		if elemento%3 == 0 {
			slice = append(slice, elemento)
		}
	}

	fmt.Println("novo slice:", slice)
}
