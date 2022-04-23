package main

import "fmt"

type Categoria struct {
	Nome string
	Pai  *Categoria
}

// type Pessoa struct {
// 	Nome      string
// 	Sobrenome string
// 	Idade     uint8
// 	Status    bool
// }

func (c Categoria) hasParent() bool {
	return c.Pai != nil
}

// func (p Pessoa) String() string {
// 	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
// }

func (c *Categoria) SetPai(pai *Categoria) {
	c.Pai = pai
}

func main() {
	p := Pessoa{"Alan", "Barcelos", 23, true}
	fmt.Println(p)

	cat := Categoria{Nome: "Categoria 1"}
	cat.SetPai(&Categoria{Nome: "Pai"})
	if !cat.hasParent() {
		fmt.Println("não tem pai!")
	}
}
