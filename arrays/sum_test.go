package arrays

import "testing"

func TestSum(t *testing.T) {
	t.Run("Coleção de tamanho 5", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		result := Sum(numbers)
		expected := 15

		if result != expected {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", result, expected, numbers)
		}
	})

	t.Run("Coleção sem tamanho definido", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		result := Sum(numbers)
		expected := 6

		if result != expected {
			t.Errorf("resultado '%d', esperado '%d', dado '%v'", result, expected, numbers)
		}
	})
}
