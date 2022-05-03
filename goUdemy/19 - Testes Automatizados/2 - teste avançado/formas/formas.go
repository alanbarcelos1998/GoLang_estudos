package formas

import (
	"math"
)

type Forma interface {
	area() float64
}

type Retangulo struct {
	altura  float64
	largura float64
}

func (r Retangulo) Area() float64 {
	return r.altura * r.largura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}
