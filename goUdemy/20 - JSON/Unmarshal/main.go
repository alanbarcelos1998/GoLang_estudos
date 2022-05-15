package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	//para que o campo não apareça, basta colocar um traço no lugar da chave do json
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	cachorroEmJSON := `{"nome":"Amarelo", "raca": "vira-lata", "idade": 6}`

	var c cachorro

	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c)
}
