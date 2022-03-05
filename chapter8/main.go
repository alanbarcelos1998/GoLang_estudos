package main

import (
	"fmt"
	// Podemos colocar apelidos no pacote que criamos, colocando na frente das aspas o apelido que queremos
	m "golang-book/chapter8/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := m.Average(xs)
	fmt.Println(avg)

	min := m.Min(xs)
	fmt.Println(min)

	max := m.Max(xs)
	fmt.Println(max)
}
