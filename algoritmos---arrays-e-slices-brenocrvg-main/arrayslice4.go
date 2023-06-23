package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("tamanho do slice: ")
	fmt.Scanln(&tamanho)

	numeros := make([]int, tamanho)

	for i := 0; i < tamanho; i++ {
		fmt.Printf("valor do elemento %d: ", i+1)
		fmt.Scanln(&numeros[i])
	}
	fmt.Println(numeros)

	soma := 0
	for _, valor := range numeros {
		soma += valor
	}
	fmt.Println("a soma é: ", soma)

}
