package main

import "fmt"

func main() {
	nomes := []string{"Tiago", "Daniel", "Maria", "Mara"}

	for k, nome := range nomes {
		fmt.Println(k, " = ", nome)
	}
}
