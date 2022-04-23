package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = -1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 100000
	fmt.Println(numero2)

	// alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	var (
		numeroReal1 float32 = 123.45
		numeroReal2 float64 = 123456.78
	)

	fmt.Println(numeroReal1, numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	char := 'A'
	fmt.Println(char)

	var texto string
	fmt.Println(texto)

	var booleano1 bool
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
