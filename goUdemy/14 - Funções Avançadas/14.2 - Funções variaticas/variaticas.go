package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	result := soma(1, 2, 3, 4, 5, 10, 19, 23, 13)
	fmt.Println(result)

	escrever("Olá Mundo", 10, 5, 8, 7, 3)
}
