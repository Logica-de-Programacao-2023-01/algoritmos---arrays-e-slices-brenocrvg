package main

import "fmt"

func main() {
	//Crie um Slice de inteiros com o tamanho 5. Em seguida, solicite ao usuário que informe um número inteiro. Adicione esse número ao Slice apenas se ele não estiver presente. Imprima o Slice resultante.
	var numeros = []int{54, 99, 61, 91, 11}
	var valor int
	fmt.Println("digite um valor a ser buscado no slice: ")
	fmt.Scan(&valor)
	presente := false
	for _, elemento := range numeros {
		if elemento == valor {
			presente = true
			break
		}
	}
	if presente {
		fmt.Println("o valor já está presente no slice")
	} else {
		numeros = append(numeros, valor)
	}
	fmt.Println(numeros)
}
