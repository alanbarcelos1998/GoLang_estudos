package main

import "fmt"

func main(){
	fmt.Print("Enter a number of Fahrenheit Temperature:\t ")
	var fahrenheitTemp float64
	fmt.Scanf("%f\n", &fahrenheitTemp)

	convertToCelsius := (fahrenheitTemp - 32) * 5/9

	fmt.Printf("The Fahrenheit Temp convert to Celsius is %v\n: ", convertToCelsius )
}