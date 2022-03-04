// Go tem um tipo embutido para erros
// podemos criar nossos próprios erros usando a função New do pacote errors

package main

import (
	"errors"
	"fmt"
)

func erroPer() {
	err := errors.New("Error personalizado")

	fmt.Println(err)
}

func main() {
	err := errors.New("error message ")

	fmt.Println(err)

	erroPer()
}
