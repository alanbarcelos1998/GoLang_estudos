package main

import (
	"fmt"
)

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da funcao f")

	fmt.Println(calculosMatematicos(10, 15))
}
