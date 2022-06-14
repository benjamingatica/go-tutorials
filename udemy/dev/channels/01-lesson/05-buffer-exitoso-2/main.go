package main

import "fmt"

func main() {
	// buffered channel (canal bufer)
	ch := make(chan int, 2)

	ch <- 42
	ch <- 43

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
