package main

import "fmt"

func main() {
	//Crie um Array de inteiros com 10 elementos. Em seguida, solicite ao usuário que informe um valor e verifique se esse valor está presente no Array. Imprima o resultado da busca.
	numeros := [10]int{29, 61, 99, 120, 34, 75, 88, 10, 6, 44}
	var numero int
	fmt.Println("digite um número: ")
	fmt.Scan(&numero)

	presente := false
	for _, valor := range numeros {
		if valor == numero {
			presente = true
			break
		}
	}
	if presente {
		fmt.Println("o numero está no array")
	} else {
		fmt.Println("não está no array")
	}
}
