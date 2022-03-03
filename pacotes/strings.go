package main

import (
	"fmt"
	"strings"
)

func main() {
	// Para procurar uma string menor em uma string maior, utiliza-se a função Contains
	// Contains(s, substr string) bool
	fmt.Println(strings.Contains("test", "es"))
	// => true

	// Para contar o número de vezes que uma string menor ocorre em uuma string maior. utiliza-se a função count
	// func Count(s, sep string) int
	fmt.Println(strings.Count("test", "t"))
	// => 2

	// Para determinar se uma string maior começa com uma string menor, utilize a função HasPrefix
	// func HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix("test", "te"))
	// => true

	// Para determinar se uma string maior termina com uma string menor, utiliza-se a função HasSuffix
	// func HasSuffix(s, suffix string) bool
	fmt.Println(strings.HasSuffix("test", "st"))
	//  => true

	// Para localizar a posição de uma string menor em uma string maior, use a função Index ( ela devolve -1 se não encontrar a string)
	// func Index(s, sep string) int
	fmt.Println(strings.Index("test", "e"))
	// => 1

	// Para unir uma lista de strings em uma única string separada por outra string (por exemplo, uma vírgula), utilize a função Join
	// func Join(a []string, sep string) string
	fmt.Println(strings.Join([]string{"a", "b"}, "-"))
	// => "a-b"

	//  Para repetir uma string, utilize a função Repeat:
	// func Repeat(s string, count int) string
	fmt.Println(strings.Repeat("a", 5))
	//  => "aaaaa"

	// Para substituir uma string menor por outra string em uma string maior, utilize a função Replace.
	// Em Go, Replace também aceita um número que indica quantas vezes a substituição deve ser feita (passe -1 para fazê-lo o máximo de vezes possível)
	// func Replace(s, old, new string, n int) string
	fmt.Println(strings.Replace("aaaa", "a", "b", 2))
	// => "bbaa"

	// Para separar uma string em uma lista de strings de acordo com uma string de separação (por exemplo, uma vírgula), utilize a função Split (Split é o inverso de Join)
	// func Split(s, sep string) []string
	fmt.Println(strings.Split("a-b-c-d-e", "-"))
	// => []string{"a","b","c","d","e"}

	// Para converter uma string em letras minúsculas, utilize a função ToLower
	// func ToLower(s string) string
	fmt.Println(strings.ToLower("TEST"))
	// => "test"

	// Para converter uma string em letras maiúsculas, utilize a função ToUpper
	// func ToUpper(strings.ToUpper("test"))
	fmt.Println(strings.ToUpper("test"))
	// "TEST"

	// Para converter uma string em uma fatia de bytes (e vice-versa)
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})

	fmt.Println(arr)
	fmt.Println(str)
}
