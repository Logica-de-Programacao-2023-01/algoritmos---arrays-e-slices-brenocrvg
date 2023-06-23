package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&n)

	if n <= 0 {
		fmt.Println("Número inválido. O número deve ser inteiro positivo.")
		return
	}

	numerosPrimos := make([]int, 0, n)
	count := 0
	numero := 2

	for count < n {
		if ehPrimo(numero) {
			numerosPrimos = append(numerosPrimos, numero)
			count++
		}
		numero++
	}

	fmt.Println("Números primos:", numerosPrimos)
}
func ehPrimo(numero int) bool {
	if numero < 2 {
		return false
	}

	limite := int(math.Sqrt(float64(numero)))

	for i := 2; i <= limite; i++ {
		if numero%i == 0 {
			return false
		}
	}

	return true
}

