package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("exercise 03")
	fmt.Println("Go routines:", runtime.NumGoroutine())
	var count int64
	fmt.Println(count)
	gr := 100
	var wg sync.WaitGroup
	wg.Add(gr)

	fmt.Println("Go routines:", runtime.NumGoroutine())

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			fmt.Println("count:", atomic.LoadInt64(&count))
			runtime.Gosched()
			wg.Done()
		}()
	}

	fmt.Println("Go routines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("final count:", count)
}
