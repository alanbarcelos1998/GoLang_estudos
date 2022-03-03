package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		// trata o erro aqui
		return
	}

	defer file.Close()

	// obtém o tamanho do arquivo
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// lê o arquivo
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}
