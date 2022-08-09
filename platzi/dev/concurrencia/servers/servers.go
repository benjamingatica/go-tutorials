package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for i < 10 {
		for _, server := range servers {
			go reviewServer(server, ch)
		}
		time.Sleep(1 * time.Second)
		fmt.Println(<-ch)
		i++
	}

	end := time.Since(start)
	fmt.Printf("Duration %s\n", end)
}

func reviewServer(server string, ch chan string) {
	_, err := http.Get(server)
	if err != nil {
		ch <- server + " server is not available =("
	} else {
		ch <- server + " server is available =)"
	}
}

/*
package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start := time.Now()

	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	wg.Add(len(servers))

	for _, server := range servers {
		go reviewServer(server)
	}

	wg.Wait()
	end := time.Since(start)
	fmt.Printf("Duration %s\n", end)
}

func reviewServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println(server, "server is not available =(")
		wg.Done()
		return
	}

	fmt.Println(server, "server is available =)")
	wg.Done()
}

*/
