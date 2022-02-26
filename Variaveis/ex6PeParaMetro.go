package main

import "fmt"

func main(){
	fmt.Printf("Numero de pés: ")
	var pe float64
	fmt.Scanf("%f\n", &pe)

	metros := pe * 0.3048

	fmt.Printf("O valor de %v para metros é: %v\n",pe,metros )
}