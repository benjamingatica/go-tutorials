// package palabra provides functions to handle string
package palabra

import "strings"

// ConteoUso take a string and return a map with string key and int values
func ConteoUso(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// Conteo takes a string and return the words amount
func Conteo(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
