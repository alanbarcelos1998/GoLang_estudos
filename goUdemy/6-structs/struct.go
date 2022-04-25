package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome, u.idade = "Alan", 23

	fmt.Println(u)

	usuario2 := usuario{
		"Alan",
		23,
		endereco{
			"Rua Domingos martins",
			91,
		},
	}

	fmt.Println(usuario2)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
