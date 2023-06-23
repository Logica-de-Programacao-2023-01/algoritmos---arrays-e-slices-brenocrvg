package main

import "fmt"

func main() {
	//Crie um Array de floats com 4 elementos e calcule o produto dos valores armazenados no Array.
	var numeros = [4]float64{4.31, 6.33, 2.04, 5.09}
	produto := 1.0
	for _, valor := range numeros {
		produto *= valor
	}
	fmt.Println("o produto dos floats é: ", produto)

}
