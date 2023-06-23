package main

import (
	"fmt"
	"math"
)

func main() {
	slice := make([]int, 10)
	slice[0] = 7845790
	slice[1] = 509
	slice[2] = 821
	slice[3] = 222
	slice[4] = 19873
	slice[5] = 7
	slice[6] = 6
	slice[7] = 10
	slice[8] = 55
	slice[9] = 61

	minimo := math.MaxInt64
	maximo := math.MinInt64

	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}

		if valor > maximo {
			maximo = valor
		}
	}

	fmt.Println("valor mínimo:", minimo)
	fmt.Println("valor máximo:", maximo)
}

