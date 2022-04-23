package main

import "fmt"

func main() {
	primary := 0
	secondary := 1
	count := 0
	// mostrarOla(0)
	fmt.Printf("%v %v ", primary, secondary)
	fibonnaci(primary, secondary, count, 5)
}

func fibonnaci(primary, secondary, count, limit int) {
	next := primary + secondary
	fmt.Printf("%v ", next)
	primary = secondary
	secondary = next
	if count >= limit {
		return
	}
	fibonnaci(primary, secondary, count+1, limit)
}
