// package matematicas provides arithmetic functions
package matematicas

import "sort"

// provides the average amount between the three central values
func PromedioCentrado(xi []int) float64 {
	sort.Ints(xi)
	xi = xi[1:(len(xi) - 1)]

	n := 0
	for _, v := range xi {
		n += v
	}

	f := float64(n) / float64(len(xi))
	return f
}
