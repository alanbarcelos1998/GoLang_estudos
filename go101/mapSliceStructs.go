package main

import "fmt"

/*func main() {
	// slice
	nomes := []string{"Tiago", "Dani", "Marcos", "Maria"}
	fmt.Println(nomes, len(nomes), cap(nomes))

	nomes = append(nomes, "Rafael")
	fmt.Println(nomes, len(nomes), cap(nomes))
}*/

/*func main() {
	// Map
	idades := make(map[string]uint8)
	idades["Tiago"] = 31
	idades["Dani"] = 36
	idades["Maria"] = 23

	val, ok := idades["Lucas"]
	fmt.Println(val, ok)
}*/

// type Pessoa struct {
// 	Nome      string
// 	Sobrenome string
// 	Idade     uint8
// 	Status    bool
// }

func main() {
	// Structs
	p := Pessoa{
		Nome:      "Alan",
		Sobrenome: "Barcelos",
		Idade:     23,
		Status:    true,
	}

	fmt.Println(p.Nome, p.Sobrenome)
}
