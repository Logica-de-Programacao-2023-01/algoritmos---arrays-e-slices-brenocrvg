package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 8)
	var valor string
	fmt.Println("Informe um valor:")
	fmt.Scanln(&valor)

	slice[0] = "andar"
	slice[1] = "seguir"
	slice[2] = "informar"
	slice[3] = "vasco"
	slice[4] = "saudade"
	slice[5] = "maravilha"
	slice[6] = "brilhante"
	slice[7] = "corrente"

	resultado := remover(slice, valor)
	fmt.Println("Slice resultante:", resultado)
}

func remover(slice []string, valor string) []string {
	resultado := []string{}

	for _, elemento := range slice {
		if elemento != valor {
			resultado = append(resultado, elemento)
		}
	}

	return resultado
}
