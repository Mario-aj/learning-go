package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	result := Perimeter(rectangle)
	expected := 40.0

	if result != expected {
		t.Errorf("result %.2f, expected %.2f", result, expected)
	}
}

func TestArea(t *testing.T) {
	verifyArea := func(t *testing.T, form Form, expected float64) {
		t.Helper()

		result := form.Area()

		if result != expected {
			t.Errorf("result %.2f, expected %.2f", result, expected)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}

		verifyArea(t, rectangle, 72.0)
	})

	t.Run("Circles", func(t *testing.T) {
		circle := Circle{10}

		verifyArea(t, circle, 314.1592653589793)
	})
}

func TestArea2(t *testing.T) {
	testArea := []struct {
		name     string
		form     Form
		expected float64
	}{
		{name: "Rectangle", form: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", form: Circle{Ray: 10}, expected: 314.1592653589793},
		{name: "Triangle", form: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range testArea {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.form.Area()

			if result != tt.expected {
				t.Errorf("%#v result %.2f, expected %.2f", tt.form, result, tt.expected)
			}
		})
	}
}
