package main

import "fmt"

func half(x int) (int, bool) {
	half := x / 2

	if half%2 == 1 {
		return 0, false
	} else {
		return 1, true
	}
}

func main() {
	fmt.Println(half(12))
}
