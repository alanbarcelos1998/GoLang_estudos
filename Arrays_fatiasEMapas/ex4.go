package main

import "fmt"

func main() {

	// var min int

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for i := 0; i < len(x); i++ {
		if x[i] < i-1 {
			fmt.Println(x[i])
		}
	}

	// for i, v := range x {
	// 	if i == 0 || v < min {
	// 		min = v
	// 	}
	// }
	// fmt.Printf("O menor valor dessa lista é: %v\n", min)

	// for i := 0; i < len(x); i++ {
	// 	if x[i] < x[i] {
	// 		fmt.Printf("O menor valor dessa lista é: %v,", x[i])
	// 	}
	// }
}
