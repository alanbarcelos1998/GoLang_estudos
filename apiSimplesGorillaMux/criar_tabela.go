package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DbCreate() {
	fmt.Println("-- Programa de Criar Tabela --")
	// abrir a conexão
	fmt.Println("-- Abrindo Conexão --")
	db, erroAbertura := sql.Open("mysql", "root:root@tcp(localhost:3307)/crudsimples")

	if erroAbertura != nil {
		log.Fatal(erroAbertura.Error())
	}

	erroPing := db.Ping()

	if erroPing != nil {
		log.Fatal(erroPing.Error())
	}

	fmt.Println("Criando tabela...")
	// criar tabela
	/*
		_, erroCreate := db.Exec(
			"CREATE TABLE livros (" +
				"id INT NOT NULL AUTO_INCREMENT," +
				"autor VARCHAR(50) NOT NULL," +
				"titulo VARCHAR(100) NOT NULL," +
				"PRIMARY KEY(id)" +
				")")

		if erroCreate != nil {
			log.Fatal(erroCreate.Error())
		}
	*/

	// criar registros
	// fmt.Println("Criando Registros")

	// _, erroInsercao := db.Exec("INSERT INTO livros (autor, titulo) VALUES ('José de Alencar', 'O Guarani')")

	// if erroInsercao != nil {
	// 	log.Fatal(erroInsercao.Error())
	// }

	fmt.Println("Pronto!")
}
