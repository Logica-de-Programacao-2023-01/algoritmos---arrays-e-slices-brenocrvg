package main

import "fmt"

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var slice []int

	for _, elemento := range array {
		if elemento%2 == 0 {
			slice = append(slice, elemento)
		}
	}

	fmt.Println("Novo Slice:", slice)
}
