package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Joao",
		sobrenome: "Da Silva",
		fumante:   false,
	}

	cliente2 := cliente{
		nome:      "Joana",
		sobrenome: "De Jesus",
		fumante:   true,
	}

	fmt.Println(cliente1)
	fmt.Println(cliente2)

}
