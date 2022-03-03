package main

import "fmt"

func parVar(x ...int) int {
	var min int

	for i, v := range x {
		if i == 0 || v < min {
			min = v
		}
	}
	return min
}

func main() {
	x := parVar(3, 5, 8, 7, 4, 56)

	fmt.Println(x)
}
