package main

import "fmt"

func main() {
	array := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	var linha int
	fmt.Println("índice de linha:")
	fmt.Scanln(&linha)

	var coluna int
	fmt.Println("índice de coluna:")
	fmt.Scanln(&coluna)

	if linha >= 0 && linha < len(array) && coluna >= 0 && coluna < len(array[linha]) {
		valor := array[linha][coluna]
		fmt.Println("valor armazenado:", valor)
	} else {
		fmt.Println("indices inválidos")
	}
}
