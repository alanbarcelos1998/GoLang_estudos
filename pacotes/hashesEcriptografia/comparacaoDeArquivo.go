// Comparacao de arquivos por hashs não criptografadas

package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// abre o arquivo
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// lembre-se de sempre fechar arquivos abertos
	defer f.Close()

	// cria um hasher
	h := crc32.NewIEEE()

	// copia o arquivo no hasher
	// - Copy aceita(dst, src) e devolve (bytesWritten, error)
	_, err = io.Copy(h, f)

	// não nos importamos com o número de bytes escritos, mas queremos

	// tratar o erro
	if err != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func main() {
	h1, err := getHash("test1.txt")

	if err != nil {
		return
	}

	h2, err := getHash("test2.txt")

	if err != nil {
		return
	}

	fmt.Println(h1, h2, h1 == h2)
}
