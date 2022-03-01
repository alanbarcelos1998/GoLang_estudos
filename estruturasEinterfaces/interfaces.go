// Uma interface é criada com a palavra reservada type, seguida de um nome e da palavra reservada interface

// Em vez de definir campos, definimos um conjunto de método (method set)

package main

import "fmt"

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64{
	var area float64
	for_ , s := range shapes{
		area += s.area()
	}
	return area
}

func main(){
	// fmt.Println(totalArea(&c, &r))

}
