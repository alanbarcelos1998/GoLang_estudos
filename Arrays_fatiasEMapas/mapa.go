// Um mapa é uma coleção não ordenada de pares chave-valor
// Mapas são usados para buscar um valor de acordo com sua chave associada
// Eis um exemplo:
// var x map[string]int
// O tipo map é representado pela palavra reservada map, seguida do tipo da chave entre colchetes e, por fim, o tipo do valor.

package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["key"] = 10
	// é possível ver o tamanho do mapa usando a função len()
	// Podemos apagar itens de um mapa usando a função embutida delete
	// exemplo: delete(x,1)
	fmt.Println(x)
}
