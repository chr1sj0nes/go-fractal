package fract

import (
	"image"
	"math"
	"testing"
)

func TestMandelbrot(t *testing.T) {
	r := image.Rect(0, 0, 300, 200)
	b := Bounds{complex(-2, -1), complex(1, 1)}
	img := NewBinaryImage(r)

	Mandelbrot(img, b)
	
	// check area of the set
	const expected = 1.50659177 
	actual := float64(img.Count()) * b.Area() / float64(r.Dx()*r.Dy())
	
	if math.Abs(expected - actual) / expected > 0.005 {
		t.Errorf("Mandelbrot Area: expected = %f, actual = %f", expected, actual)
	}
}
