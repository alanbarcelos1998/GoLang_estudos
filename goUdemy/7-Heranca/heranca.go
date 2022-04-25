package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Jose", "Silva", 50, 1.65}

	e1 := estudante{p1, "Direito", "USP"}
	fmt.Println(e1)
	fmt.Println(e1.nome)
}
