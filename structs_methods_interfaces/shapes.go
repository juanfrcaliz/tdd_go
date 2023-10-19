package structsmethodsinterfaces

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

// Rectangle related definitions
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle related definitions
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle related definitions
type Triangle struct {
	l1 float64
	l2 float64
	l3 float64
}

func (t Triangle) Area() float64 {
	return 0.25 * math.Sqrt(4*math.Pow(t.l1, 2)*math.Pow(t.l2, 2)-math.Pow(math.Pow(t.l1, 2)+math.Pow(t.l2, 2)-math.Pow(t.l3, 2), 2))
}

func (t Triangle) Perimeter() float64 {
	return t.l1 + t.l2 + t.l3
}
