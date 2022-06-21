package perrito

import (
	"fmt"
	"testing"
)

var hAge = 26
var expResp = 182

func TestEdad(t *testing.T) {
	r := Edad(hAge)

	if r != expResp {
		t.Error("Expected", expResp, "GOT", r)
	}
}

func ExampleEdad() {
	fmt.Println(Edad(hAge))
	// Output:
	// 182
}

func BenchmarkEdad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Edad(hAge)
	}
}

func TestEdadDos(t *testing.T) {
	r := EdadDos(hAge)

	if r != expResp {
		t.Error("Expected", expResp, "GOT", r)
	}
}

func ExampleEdadDos() {
	fmt.Println(EdadDos(hAge))
	// Output
	// 182
}

func BenchmarkEdadDos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EdadDos(hAge)
	}
}
