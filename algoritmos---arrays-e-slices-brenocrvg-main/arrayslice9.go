package main

import (
	"fmt"
)

func main() {
	floats := [6]float64{}
	var numero float64
	fmt.Println("Informe um número:")
	fmt.Scanln(&numero)

	for i := 0; i < len(floats); i++ {
		floats[i] = numero
	}

	fmt.Println("array resultante:", floats)
}
