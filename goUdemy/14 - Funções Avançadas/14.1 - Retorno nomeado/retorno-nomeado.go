package main

import "fmt"

func calcMatematicos(n1, n2 int) (soma int, sub int) {
	soma = n1 + n2
	sub = n1 - n2
	return
}

func main() {
	n1 := 10
	n2 := 20

	fmt.Println(calcMatematicos(n1, n2))
}
