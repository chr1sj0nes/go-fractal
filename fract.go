// Functions for generating images of fractal sets
package fract

import (
	"image"
	"image/color"
	"math/cmplx"
)

const MaxIterations = 10000

// Fractal image receiver
type Image interface {
	SetPixel(x, y, iterations int)
	Bounds() image.Rectangle
}

// Generate Mandelbrot set
func Mandelbrot(img Image, min, max complex128) {
	b := img.Bounds()
	dr := real(max-min) / float64(b.Dx())
	di := imag(max-min) / float64(b.Dy())

	ch := make(chan bool, b.Dx())

	for x := 0; x < b.Dx(); x++ {
		go func(x int) {
			for y := 0; y < b.Dy(); y++ {
				z := complex(0, 0)
				c := min + complex(float64(x)*dr, float64(y)*di)

				n := 0
				for ; n < MaxIterations; n++ {
					z = z*z + c
					if cmplx.Abs(z) > 2.0 {
						break // we've diverged
					}
				}

				img.SetPixel(x+b.Min.X, y+b.Min.Y, n)
			}

			ch <- true
		}(x)
	}

	// wait for all go routines to finish
	for i := 0; i < b.Dx(); i++ {
		<-ch
	}
}

type ColorImage struct {
	image.RGBA
}

func (img *ColorImage) SetPixel(x, y, iterations int) {
	color := color.RGBA{128, 0, 0, 0} // TODO
	img.Set(x, y, color)
}

type BinaryImage struct {
	Pix  []bool
	Rect image.Rectangle
}

func NewBinaryImage(r image.Rectangle) *BinaryImage {
	return &BinaryImage{make([]bool, r.Dx()*r.Dy()), r}
}

func (img *BinaryImage) SetPixel(x, y, iterations int) {
	n := x - img.Bounds().Min.X + (y-img.Bounds().Min.Y)*img.Bounds().Dx()
	img.Pix[n] = (iterations == MaxIterations)
}

func (img BinaryImage) Bounds() image.Rectangle {
	return img.Rect
}

func (img BinaryImage) Count() uint {
	count := uint(0)
	for _, pix := range img.Pix {
		if pix {
			count++
		}
	}

	return count
}
