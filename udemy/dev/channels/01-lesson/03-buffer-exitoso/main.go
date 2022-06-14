package main

import "fmt"

func main() {
	// buffered channel (canal bufer)
	ch := make(chan int, 1)

	ch <- 42

	fmt.Println(<-ch)
}
