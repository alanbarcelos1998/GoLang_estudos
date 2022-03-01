// Uma estrutura é um tipo que contém campos nomeados

package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// estrutura
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

// MÉTODO
// entre a palavra reservada func e o nome da função, adicionamos um receptor
// O receptor é como um parâmetro, tem um nome e um tipo
// ao criar dessa maneira, podemos chamar a função usando .
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func main() {
	// Os campos da struct começam de acordo com o valor padrão de cada tipo que foi atribuído
	// 0 para int, 0.0 para float, "" para string, nil para ponteiros

	// para atribuirmos valores nos campos da struct podemos fazer
	c := Circle{0, 0, 5}
	// Se quiser um ponteiro para struct, basta fazer
	// 	c := &Circle{0,0,5}

	r := Rectangle{0, 0, 10, 10}

	// Podemos acessar os campos usando o operador .
	// fmt.Println(c.x, c.y, c.r)
	// c.x = 10
	// c.y = 5

	fmt.Println(c.area())
	fmt.Println(r.area())
}
