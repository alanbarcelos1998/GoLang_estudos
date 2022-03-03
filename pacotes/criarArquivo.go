package main

import "os"

func main() {
	file, err := os.Create("teste2.txt")
	if err != nil {
		// trata o erro aqui
		return
	}

	defer file.Close()

	file.WriteString("test")
}
