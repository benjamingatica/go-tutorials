package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer fmt.Println("hello world")
	_, err := os.Open("no-file.txt")
	if err != nil {
		// log.Println("encountered error: ", err) // run defer code
		// return

		// panic(err) // run defer code

		// fmt.Println(err) // run defer code
		// return

		// log.Fatalln(err) // does not run defer code

		log.Panicln(err) // run defer code
	}

	fmt.Println("hello world 2")
}
