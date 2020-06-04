package main

import (
	"fmt"
)

func isNumberEven(number int) {
	if number%2 == 0 {
		fmt.Printf("Число %d - чётное\n", number)
	} else {
		fmt.Printf("Число %d - нечётное\n", number)
	}
}

func isNumberDividedByThree(number int) {
	if number%3 == 0 {
		fmt.Printf("Число %d делится на 3 без остатка\n", number)
	} else {
		fmt.Printf("Число %d не делится на 3 без остатка\n", number)
	}
}

func Fibonachi(n int) int {

	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	return Fibonachi(n-1) + Fibonachi(n-2)
}

func FibonachiLsit() {
	for i := 0; i < 100; i++ {
		fmt.Println(i+1, Fibonachi(i))
	}
}

func main() {
	isNumberEven(5)
	isNumberDividedByThree(6)
	FibonachiLsit()
}
