package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesArrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipoDeConstrucao string
	tamanhoDaLoucura string
}

func (x dentista) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, " e eu já arranquei", x.dentesArrancados, "dentes e ouve só: Bom dia!")
}

func (x arquiteto) oiBomDia() {
	fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oiBomDia()
}

func serHumano(g gente) {
	g.oiBomDia()
}

func main() {
	mrDente := dentista{
		pessoa: pessoa{
			nome:      "Mister Dente",
			sobrenome: "Borges",
			idade:     50,
		},
		dentesArrancados: 5000,
		salario:          25000.00,
	}

	mrPredio := arquiteto{
		pessoa: pessoa{
			nome:      "Mister Arquiteto",
			sobrenome: "Da Silva",
			idade:     45,
		},
		tipoDeConstrucao: "Arranha ceu",
		tamanhoDaLoucura: "Alta",
	}

	mrDente.oiBomDia()
	mrPredio.oiBomDia()
}
