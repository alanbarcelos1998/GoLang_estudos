package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

type Profissional struct {
	pessoa  Pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := Pessoa{
		nome:  "Alfredo",
		idade: 30,
	}

	pessoa2 := Profissional{
		pessoa: Pessoa{
			nome:  "Maricota",
			idade: 31,
		},
		titulo:  "Designer",
		salario: 1850,
	}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
}
