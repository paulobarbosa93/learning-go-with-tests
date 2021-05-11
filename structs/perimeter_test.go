package structs

import "testing"

func TestPerimeter(t *testing.T) {
	retangle := Retangle{10.0, 10.0}
	result := Perimeter(retangle)
	expected := 40.0

	if result != expected {
		t.Errorf("resultado %2.f, esperado %2.f", result, expected)
	}
}

func TestArea(t *testing.T) {
	// "table driven tests" (testes orientados por tabela)
	testsArea := []struct {
		name     string
		form     Form
		expected float64
	}{
		{name: "Retangle", form: Retangle{Width: 12.0, Height: 6.0}, expected: 72.0},
		{name: "Circle", form: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range testsArea {
		t.Run(tt.name, func(t *testing.T) {
			resultado := tt.form.Area()
			if resultado != tt.expected {
				t.Errorf("%#v resultado %.2f, esperado %.2f", tt, resultado, tt.expected)
			}
		})
	}
	// "table driven tests" (testes orientados por tabela)

	checkArea := func(t *testing.T, form Form, expected float64) {
		t.Helper()

		result := form.Area()

		if result != expected {
			t.Errorf("resultado %2.f, esperado %2.f", result, expected)
		}
	}

	t.Run("Retângulos", func(t *testing.T) {
		retangle := Retangle{12.0, 6.0}
		expected := 72.0

		checkArea(t, retangle, expected)
	})

	t.Run("Círculos", func(t *testing.T) {
		circle := Circle{10}
		expected := 314.1592653589793

		checkArea(t, circle, expected)
	})
}
