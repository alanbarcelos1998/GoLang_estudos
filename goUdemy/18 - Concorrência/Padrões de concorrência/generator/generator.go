package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("Ola Mundo!")

	i := 0
	for i <= 10 {
		fmt.Println(<-canal)
		i++
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
