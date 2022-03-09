// canais oferecem um mode de duas goroutines se comunicarem
package main

import (
	"fmt"
	"time"
)

// Um tipo canal é representado pela palavra chan
// Quando a seta está depois do chan, a função só enviará dados
func pinger(c chan<- string) {
	for i := 0; ; i++ {
		//  a <- é usado para enviar e receber mensagem pelo canal
		c <- "ping"
	}
}

// Quando a seta está antes de chan, ele somente recebe dados
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

// Existem também canais com Buffer, por exemplo:
// c := make(chan int, 20)
// Um canal com buffer é assíncrono
// se o canal tiver cheio provocará uma espera até que haja espaço para pelo menos
// mais um int
