package main

import (
	"fmt"
)

func main() {
	fmt.Println("The program calculates the resistance R of two resistors.")
	fmt.Println("Enter the resistance of two resistors:")
	var a, b float64
	fmt.Scanln(&a, &b)
	R := a + b
	G := 1/a + 1/b
	R3 := 1 / G 
	fmt.Println("Resistance when connected in series is", R, "; resistance when connected in parallel is", R3)
}