package main

import "fmt"

type Document interface {
	Doc() string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu CNPJ é: %s", pj.cnpj)
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	pf := PessoaFisica{
		Pessoa{Nome: "Alan", Idade: 23, Status: true},
		"Barcelos",
		"162.528.107-85",
	}

	show(pf)

	pj := PessoaJuridica{
		Pessoa{Nome: "Jose", Idade: 5, Status: true},
		"Jose LTDA",
		"10.025.065/0001-85",
	}

	show(pj)
}
