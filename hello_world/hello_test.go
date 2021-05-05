package hello_world

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Paul")
	expected := "Hello, Paul"

	if result != expected {
		t.Errorf("resultado %q, esperado %q", result, expected)
	}
}
