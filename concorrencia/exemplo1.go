// goroutine é uma função capaz de executar de modo concorrente com outras funções
// Para criar uma goroutine utilizamos a palavra reservada go, seguida de uma chamada de função

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(0)
	}
	var input string
	fmt.Scanln(&input)
}
