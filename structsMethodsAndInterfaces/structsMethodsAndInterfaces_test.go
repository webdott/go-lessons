package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f hasPerimeter %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	assertArea := func(t testing.TB, shape struct {
		name  string
		shape Shape
		want  float64
	}) {
		t.Helper()
		got := shape.shape.Area()
		if got != shape.want {
			t.Errorf("%#v got %g hasArea %g", shape.shape, got, shape.want)
		}
	}

	areasTest := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, want: 72.0},
		{name: "Circle", shape: Circle{10}, want: 314.1592653589793},
		{name: "Triangle", shape: Triangle{12, 6}, want: 36.0},
	}

	for _, areaTest := range areasTest {
		t.Run(areaTest.name, func(t *testing.T) {
			assertArea(t, areaTest)
		})
	}
}
