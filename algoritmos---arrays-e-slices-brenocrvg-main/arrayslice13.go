package main

import "fmt"

func main() {
	array := [7]int{}
	var numero int
	fmt.Println("número:")
	fmt.Scanln(&numero)
	array[0] = numero
	array[len(array)-1] = numero
	fmt.Println("array novo:", array)
}
