package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
		fmt.Println(cara)
	case "coroa":
		coroa++
		fmt.Println(coroa)
	default:
		fmt.Println("caiu em pé")
	}
}

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)
	if _, erro := file.Read(data); erro != nil {
		log.Panic(erro)
	}

	fmt.Println(string(data))

	lancarMoeda("coroa")
	lancarMoeda("coroa")
}
