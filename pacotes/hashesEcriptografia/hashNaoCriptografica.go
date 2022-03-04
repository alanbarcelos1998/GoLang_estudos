package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// cria um hasher
	h := crc32.NewIEEE()
	// escreve nossos dados no hasher
	h.Write([]byte("test"))
	// calcula o checksum crc32
	v := h.Sum32()
	fmt.Println(v)
}
