package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	result := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if result != expected {
		t.Errorf("esperado %q, recebido %q", expected, result)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("a", 10)
	fmt.Println(result)
	// Output: aaaaaaaaaa
}
