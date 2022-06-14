package main

import "fmt"

func main() {
	// unbuffered channel (canal sin bufer)
	ch := make(chan int)

	go func() {
		ch <- 42
	}()

	fmt.Println(<-ch)
}
