package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("The program finds solutions of a quadratic equation.")
	fmt.Println("Enter the coefficients a, b and c:")
	var a, b, c float64
	fmt.Scanln(&a, &b, &c)
	D := b*b - 4*a*c 
	x1 := (-b + math.Sqrt(D)) / (2 * a)
	x2 := (-b - math.Sqrt(D)) / (2 * a)
	fmt.Println("x1=", x1)
	fmt.Println("x2=", x2)
}