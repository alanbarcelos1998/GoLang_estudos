package main

import "fmt"

func fatorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

func fatorialComLoop(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}

func main() {
	fmt.Println(fatorial(3))
	fmt.Println(fatorialComLoop(3))
}
