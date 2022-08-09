package main

import (
	"fmt"
	"log"
)

func main() {
	/* scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	operation := scanner.Text()
	fmt.Println(operation)
	values := strings.Split(operation, "+")
	fmt.Println(values)
	fmt.Println(values[0] + values[1])
	operator1, err := strconv.Atoi(values[0])
	if err != nil {
		log.Fatalln(err)
	}
	operator2, err := strconv.Atoi(values[1])
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(operator1 + operator2) */
	var a, b int
	var operator string
	fmt.Println("enter first value: ")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("enter second value: ")
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("enter operator: ")
	_, err = fmt.Scan(&operator)
	if err != nil {
		log.Fatalln(err)
	}

	switch operator {
	case "+":
		result(sum(a, b))
	case "-":
		result(subtraction(a, b))
	case "*":
		result(multiplication(a, b))
	case "/":
		result(division(a, b))
	default:
		fmt.Println("case not supported")
	}

}

func result(x int) {
	fmt.Println("result:", x)
}

func sum(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}

func division(a, b int) int {
	return a / b
}
