package matematicas

import (
	"fmt"
	"testing"
)

var expRes = 6.0

func TestPromedioCentrado(t *testing.T) {
	r := PromedioCentrado([]int{1, 4, 6, 8, 100})

	if r != expRes {
		t.Error("Expected", expRes, "GOT", r)
	}
}

func ExamplePromedioCentrado() {
	fmt.Println(PromedioCentrado([]int{1, 4, 6, 8, 100}))
	// Output:
	// 6
}

func BenchmarkPromedioCentrado(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PromedioCentrado([]int{1, 4, 6, 8, 100})
	}
}
