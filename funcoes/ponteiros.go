package main

import "fmt"

func addPont(y int) int {
	return y
}

var y int = 5

func main() {
	x := &y

	fmt.Println(addPont(*x))
}
