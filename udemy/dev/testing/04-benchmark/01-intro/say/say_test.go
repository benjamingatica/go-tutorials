package say

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet("Benja")
	expResp := "Welcome dear Benja"
	if s != expResp {
		t.Error("Expected", expResp, "Got", s)
	}
}

func ExampleGreet() {
	fmt.Println(Greet("Benja"))
	// Output:
	// Welcome dear Benja
}

func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Benja")
	}
}
