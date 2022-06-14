package main

import "fmt"

func main() {
	// unbuffered channel (canal sin bufer)
	ch := make(chan int)
	ch <- 42
	fmt.Println(<-ch)
}
