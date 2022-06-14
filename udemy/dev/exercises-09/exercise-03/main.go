package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("exercise 03")
	fmt.Println("Go routines:", runtime.NumGoroutine())
	count := 0
	fmt.Println(count)
	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	fmt.Println("Go routines:", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {
			code := count
			runtime.Gosched()
			code++
			count = code
			fmt.Println("count:", count)
			wg.Done()
		}()
	}

	fmt.Println("Go routines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("final count:", count)
}
