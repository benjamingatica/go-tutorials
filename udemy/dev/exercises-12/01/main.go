package main

import (
	"01/dog"
	"fmt"
)

func main() {
	hAge := 26
	dAge := dog.Age(hAge)
	fmt.Println(dAge)
}
