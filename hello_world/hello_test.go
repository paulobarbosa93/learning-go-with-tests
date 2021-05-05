package hello_world

import (
	"testing"
)

func TestHello(t *testing.T) {
	checkCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("resultado %q, esperado %q", result, expected)
		}
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		result := Hello("Paul", "")
		expected := "Hello, Paul"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia por passada", func(t *testing.T) {
		result := Hello("", "")
		expected := "Hello, world"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("em espanhol", func(t *testing.T) {
		result := Hello("Paul", "espanhol")
		expected := "Hola, Paul"
		checkCorrectMessage(t, result, expected)
	})

	t.Run("em francês", func(t *testing.T) {
		result := Hello("Paul", "frances")
		expected := "Bonjour, Paul"
		checkCorrectMessage(t, result, expected)
	})
}
