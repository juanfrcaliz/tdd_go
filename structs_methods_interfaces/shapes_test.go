// We need some geometry code to calculate the perimeter of a rectangle given a height and width.
package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{11.0, 9.0}, hasPerimeter: 40.0},
		{name: "Circle", shape: Circle{5.0}, hasPerimeter: 31.41592653589793},
		{name: "Triangle", shape: Triangle{3.0, 2.0, 5.0}, hasPerimeter: 10.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("From %#v expected %g but got %g", tt.shape, tt.hasPerimeter, got)
			}
		})

	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{5.0, 4.0}, hasArea: 20.0},
		{name: "Circle", shape: Circle{7.0}, hasArea: 153.93804002589985},
		{name: "Triangle", shape: Triangle{4.0, 5.0, 2.0}, hasArea: 3.799671038392666},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("From %#v expected %g but got %g", tt.shape, tt.hasArea, got)
			}
		})
	}
}
