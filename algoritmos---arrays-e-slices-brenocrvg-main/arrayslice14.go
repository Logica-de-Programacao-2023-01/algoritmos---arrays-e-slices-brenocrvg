package main

import "fmt"

func main() {
	var slice = []int{55, 61, 90, 444, 21, 55, 22, 12}
	fmt.Println(slice)
	var indice1, indice2 int
	fmt.Println("primeiro índice (entre 0 e 7):")
	fmt.Scanln(&indice1)
	fmt.Println("segundo índice (entre 0 e 7):")
	fmt.Scanln(&indice2)

	if indice1 >= 0 && indice1 < len(slice) && indice2 >= 0 && indice2 < len(slice) {

		slice[indice1], slice[indice2] = slice[indice2], slice[indice1]
		fmt.Println("slice resultante:", slice)
	} else {
		fmt.Println("indices inválidos")
	}
}
