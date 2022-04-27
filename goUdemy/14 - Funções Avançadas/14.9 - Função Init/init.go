package main

import "fmt"

var n int

func main() {
	fmt.Println("Executando a função main")
	n = 10
}

// A função init executa antes da main
func init() {
	fmt.Println("Executando a função init")
	fmt.Println(n)
}
