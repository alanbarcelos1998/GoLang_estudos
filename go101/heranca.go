// package main

// import "fmt"

// type Pessoa struct {
// 	Nome   string
// 	Idade  uint8
// 	Status bool
// }

// type PessoaFisica struct {
// 	Pessoa
// 	Sobrenome string
// 	cpf       string
// }

// func (p PessoaFisica) String() string {
// 	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
// }

// type PessoaJuridica struct {
// 	RazaoSocial string
// 	cnpj        string
// }

// func main() {
// 	p := PessoaFisica{
// 		Pessoa{Nome: "Alan", Idade: 23, Status: true},
// 		"Barcelos",
// 		"162.528.107-85",
// 	}

// 	fmt.Println(p)
// }
