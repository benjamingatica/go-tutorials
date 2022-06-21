// package mate help to check that you know how to sum
package mate

// this function sum an unlimited amount of integer numbers
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}

	return s
}
