package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo()
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("err happened", err)
		//		log.Println("err happened", err)
		//		log.Fatalln(err)
		log.Panicln(err)
		//		panic(err)
	}

	fmt.Println("hello world2 !")
}

func foo() {
	fmt.Println("hello world!")
}

/*
Panicln es quivalente a Println() seguido por una llamada a panic().
*/

// Fatalln es equivalente a Println() seguido por una llamada a os.Exit(1).
