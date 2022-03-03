// outra maneira de obter um ponteiro é usando a função embutida new

package main

import "fmt"

func one(xPtr *int) {
	*xPtr = 3
}

func main() {
	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr)
}
