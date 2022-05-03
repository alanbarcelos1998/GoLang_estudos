package main

import (
	"fmt"
	"time"
)

// Canal pode enviar e receber dados
func main() {
	canal := make(chan string)
	go escrever("ola mundo", canal)

	fmt.Println("Depois da função escrever começar a ser executada")

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	time.Sleep(time.Second)
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
