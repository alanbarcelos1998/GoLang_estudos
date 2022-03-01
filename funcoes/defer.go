// Defer escalona uma chamada de função para ser executada após a função terminar

package main

import "fmt"

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	defer first()
	second()
}
