package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Hello world")
	fmt.Println("GO routines:", runtime.NumGoroutine())
	fmt.Println("Num CPUs:", runtime.NumCPU())

	wg.Add(2)

	fmt.Println("greetings from main")
	go foo()
	go bar()

	fmt.Println("GO routines:", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	fmt.Println("greetings from foo")
	wg.Done()
}

func bar() {
	fmt.Println("greetings from bar")
	wg.Done()
}
