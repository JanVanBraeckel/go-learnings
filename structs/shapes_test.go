package structs

import "testing"

type Shape interface {
	Area() float64
}

func TestPerimeter(t *testing.T) {
	rect := Rectangle{Width: 12.1, Height: 13.5}
	got := Perimeter(rect)
	expected := 51.2

	if got != expected {
		t.Errorf("expected %g got %g", expected, got)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
