package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Livro struct {
	Id     int    `json:"id"`
	Titulo string `json:"titulo"`
	Autor  string `json:"autor"`
	Ano    int    `json:"ano"`
}

var Livros []Livro = []Livro{
	Livro{
		Id:     1,
		Titulo: "O Guarani",
		Autor:  "Jose de Alencar",
	},
	Livro{
		Id:     2,
		Titulo: "Cazuza",
		Autor:  "Viriato Correia",
	},
	Livro{
		Id:     3,
		Titulo: "Dom Casmurro",
		Autor:  "Machado de Assis",
	},
}

func erroPersonalizado(err string) error {

	convertErro := errors.New("Erro de conversão")
	padrao := errors.New("Erro padrão")

	switch err {
	case "convertErr":
		return convertErro

	default:
		return padrao

	}

}

func rotaPrincipal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem-vindo")
}

func listarLivros(w http.ResponseWriter, r *http.Request) {

	// Encapsula o response para json em uma variavel
	encoder := json.NewEncoder(w)
	// Mostra esse json utilizando a variável Livros
	encoder.Encode(Livros)
}

func cadastrarLivro(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusCreated)

	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		// lidar com o erro
		fmt.Println("Erro ao inserir informação!")
	}

	var novoLivro Livro
	json.Unmarshal(body, &novoLivro)
	novoLivro.Id = len(Livros) + 1
	Livros = append(Livros, novoLivro)

	encoder := json.NewEncoder(w)
	encoder.Encode(novoLivro)
}

func excluirLivro(w http.ResponseWriter, r *http.Request) {
	// DELETE
	partes := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(partes[2])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		erroPersonalizado("convertErr")
	}

	indiceDoLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceDoLivro = indice
			break
		}
	}

	if indiceDoLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	ladoEsquerdo := Livros[0:indiceDoLivro]
	ladoDireito := Livros[indiceDoLivro+1 : len(Livros)]
	Livros = append(ladoEsquerdo, ladoDireito...)

	w.WriteHeader(http.StatusNoContent)
}

func modificarLivro(w http.ResponseWriter, r *http.Request) {
	partes := strings.Split(r.URL.Path, "/")
	id, erro := strconv.Atoi(partes[2])

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	corpo, erroCorpo := ioutil.ReadAll(r.Body)

	if erroCorpo != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var livroModificado Livro
	erroJson := json.Unmarshal(corpo, &livroModificado)

	if erroJson != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	indiceDoLivro := -1
	for indice, livro := range Livros {
		if livro.Id == id {
			indiceDoLivro = indice
			break
		}
	}

	if indiceDoLivro < 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	Livros[indiceDoLivro] = livroModificado

	json.NewEncoder(w).Encode(livroModificado)
}

func rotearLivros(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	partes := strings.Split(r.URL.Path, "/")

	if len(partes) == 2 || len(partes) == 3 && partes[2] == "" {
		if r.Method == "GET" {
			listarLivros(w, r)
		} else if r.Method == "POST" {
			cadastrarLivro(w, r)
		}
	} else if len(partes) == 3 || len(partes) == 4 && partes[3] == "" {
		if r.Method == "GET" {
			buscarLivro(w, r)

		} else if r.Method == "DELETE" {
			excluirLivro(w, r)
		} else if r.Method == "PUT" {
			modificarLivro(w, r)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func buscarLivro(w http.ResponseWriter, r *http.Request) {

	partes := strings.Split(r.URL.Path, "/")

	// if len(partes) > 3 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	return
	// }

	id, err := strconv.Atoi(partes[2])

	if err != nil {
		erroPersonalizado("convertErr")
	}

	for _, livro := range Livros {
		if livro.Id == id {
			json.NewEncoder(w).Encode(livro)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func configurarRotas() {
	// Rota
	http.HandleFunc("/", rotaPrincipal)
	http.HandleFunc("/livros", rotearLivros)
	http.HandleFunc("/livros/", rotearLivros)
}

func configurarServidor() {
	configurarRotas()

	fmt.Println("Servidor esta rodando na porta 1337")
	// rodar servidor
	log.Fatal(http.ListenAndServe(":9090", nil)) // DefaultServerMux
}

func main() {
	configurarServidor()

}
