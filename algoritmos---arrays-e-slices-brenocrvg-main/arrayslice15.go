package main

import "fmt"

func main() {
	array := [10]float64{21.8, 18.5, 42.32, 26.5, 1.8, 1.3, 5.2, 4.9, 2.41, 3.9}

	var slice []float64

	for _, elemento := range array {
		if elemento > 5 {
			slice = append(slice, elemento)
		}
	}
	
	fmt.Println("Novo Slice:", slice)
}
