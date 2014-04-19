package fract

import (
	"image"
	"math"
	"testing"
)

func TestMandelbrot(t *testing.T) {
	r := image.Rect(0, 0, 300, 200)
	min := complex(-2, -1)
	max := complex(1, 1)
	img := NewBinaryImage(r)

	Mandelbrot(img, min, max)

	// check area of the set
	const expected = 1.50659177
	actual := float64(img.Count()) * real(max-min) * imag(max-min) / float64(r.Dx()*r.Dy())

	if math.Abs(expected-actual)/expected > 0.005 {
		t.Errorf("Mandelbrot Area: expected = %f, actual = %f", expected, actual)
	}
}
