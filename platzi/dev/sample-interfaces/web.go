package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type WebWriter struct{}

func (WebWriter) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%T:\n", resp.Body)

	w := WebWriter{}
	io.Copy(w, resp.Body)
}

/* package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	w, err := io.Copy(os.Stdout, r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(w)

}
*/
