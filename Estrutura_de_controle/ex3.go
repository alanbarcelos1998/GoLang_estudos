package main

import "fmt"

func main() {

	for i := 1; i < 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%v: FizzBuzz\n", i)
		} else if i%3 == 0{
			fmt.Printf("%v: Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%v: Buzz\n",i)
		}
	}
}