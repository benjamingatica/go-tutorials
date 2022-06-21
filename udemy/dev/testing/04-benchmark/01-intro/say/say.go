// package say allows us to greet a person
package say

import "fmt"

// Greet function allows us greet a person
func Greet(s string) string {
	return fmt.Sprint("Welcome dear ", s)
}
