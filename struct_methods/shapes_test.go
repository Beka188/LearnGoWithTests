package struct_methods

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{width: 10.0, height: 20.0}
	got := rectangle.Perimeter()
	want := 60.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{width: 10.0, height: 20.0}, want: 200.0},
		{name: "Circle", shape: Circle{radius: 10.0}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{a: 3.0, b: 4.0, c: 5.0}, want: 6.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.want {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
			}
		})
	}
}
