// the package main provides functions to show the use of the packages and documentation
package main

import (
	"fmt"
	"sample-documentation/dog"
)

// the main function call functions from the package dog
func main() {
	hAge := 7
	dAge := dog.Age(hAge)
	fmt.Println(hAge, "human years are equivalent to ", dAge, " dog years")
}
