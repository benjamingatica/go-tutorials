package main

import "fmt"

func main() {
	// canal receive-only
	ch := make(<-chan int, 2)

	ch <- 42
	ch <- 43

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("----------------")
	fmt.Printf("%T\n", ch)
}
