package main

import "fmt"

func main() {
	array := make([]int, 5)
	fmt.Println("Digite 5 números inteiros:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Elemento %d: ", i+1)
		fmt.Scanln(&array[i])
	}
	ordenado := true
	for i := 1; i < 5; i++ {
		if array[i] < array[i-1] {
			ordenado = false
			break
		}
	}
	if ordenado {
		fmt.Println("O array está ordenado em ordem crescente.")
	} else {
		fmt.Println("O array não está ordenado em ordem crescente.")
	}
}
