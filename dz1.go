package main

import (
	"fmt"
	"os"
)

const kursdoll float64 = 70.14

func main() {
	var rubles float64
	var dollars float64
	fmt.Print("Введите сумму в рублях: ")
	fmt.Fscan(os.Stdin, &rubles)
	dollars = rubles / kursdoll
	fmt.Printf("%0.2f рублей = %0.2f долларов.\n", rubles, dollars)
}
