package main

import "fmt"

func main() {
	x := 11

	// estaRecebeOValor(x)
	estaRecebeUmPonteiro(&x)

	fmt.Println(x)
}

// faz uma cópia da variável original e incrementa um valor
func estaRecebeOValor(x int) {
	x++
	fmt.Println("Na função:", x)
}

// muda o valor da variável original, pois aponta o endereço dela
func estaRecebeUmPonteiro(x *int) {
	*x++
	fmt.Println("Na função:", *x)
}
