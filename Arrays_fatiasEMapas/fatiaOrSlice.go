//  Uma fatia é um segmento de um array. Assim como os arrays, as fatias podem ser indexadas e têm um tamanho.
// De modo diferente dos arrays, o tamanho pode mudar.
// Eis um exemplo de uma fatia:
// 	var x [] float64
//

package main

import "fmt"

func main() {
	/* Para criar uma fatia, deve-se usar a função make*/
	// x := make([]float64, 5)

	/*Outra maneira de criar fatias é por meio da expressão [low : high]*/
	// arr := [5]float64{1,2,3,4,5}
	// x := arr[0:5]

	/* para adicionar elementos no final de uma fatia é usado o append */
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	fmt.Println("---------------------------------------------------")

	/* copy aceita dois argumentos: dst e src. */
	// Todas as entradas de src são copiadas para dst
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println(slice3, slice4)
}
