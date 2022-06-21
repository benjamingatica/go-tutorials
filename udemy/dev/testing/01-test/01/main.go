package main

import "fmt"

func main() {
	fmt.Println("La suma de 2 y 4 es :", mySum(2, 4))
	fmt.Println("La suma de 1 y 5 es :", mySum(1, 5))
	fmt.Println("La suma de 6 y 7 es :", mySum(6, 7))
}

func mySum(xi ...int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}

	return sum
}
