package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	// ouve em uma porta
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// aceita uma conexão
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		// trata a conexão
		go handleServerConnection(c)
	}
}

func handleServerConnection(c net.Conn) {
	// recebe a mensagem
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	// conecta-se ao servidor
	c, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	// envia a mensagem
	msg := "Hello, World"
	fmt.Println("Sending", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
