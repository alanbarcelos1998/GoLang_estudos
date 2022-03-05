// Fazendo parse de argumentos da linha de comando

package main

import (
	"flag"
	"fmt"
	"math/rand"
)

func main() {
	// defina a flag
	maxp := flag.Int("max", 6, "the max value")
	// Faz parse
	flag.Parse()
	// Gera um número entre 0 e max
	fmt.Println(rand.Intn(*maxp))
}

// para usar podemos fazer go run parseArgs -max=100
