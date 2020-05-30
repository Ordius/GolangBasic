package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64 = 3, 4
	var c float64
	var s float64
	var p float64

	s = 0.5 * a * b
	c = math.Sqrt(a*a + b*b)
	p = a + b + c
	fmt.Printf(" Площадь треугольника = %0.1f\n Гипотенуза треугольника = %0.1f\n Периметр треугольника = %0.1f\n", s, c, p)
}
