package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("init program")
	_, err := os.Open("sin-archivo.txt")
	if err != nil {
		//		fmt.Println("un error ocurrió", err)
		//		log.Println("un error ocurrió", err)
		//		log.Fatalln(err)
		//		log.Panicln(err)
		panic(err)
	}

	fmt.Println("final")
}

// http://godoc.org/builtin#panic
