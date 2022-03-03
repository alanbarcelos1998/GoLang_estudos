package main

import "fmt"

func troca(x, y *int) {
	*x, *y = *y, *x
	fmt.Println(*x, *y)
}

func main() {
	x := 1
	y := 2

	troca(&x, &y)
}
