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

	ch := make(chan bool, b.Dx()*b.Dy())

	for x := 0; x < b.Dx(); x++ {
		for y := 0; y < b.Dy(); y++ {
			go func() {
				z := complex(0, 0)
				c := min + complex(float64(x)*dr, float64(y)*di)

				n := 0
				for ; n < MaxIterations; n++ {
					z = z*z + c
					if cmplx.Abs(z) > 2.0 {
						break
					}
				}

				img.SetPixel(x+b.Min.X, y+b.Min.Y, n)
				ch <- true
			}()
		}
	}

	// wait for all go routines to finish
	for i := 0; i < b.Dx()*b.Dy(); i++ {
		<-ch
	}
}

type ColorImage struct {
	rgb image.RGBA
}

func (img *ColorImage) SetPixel(x, y, iterations int) {
	color := color.RGBA{128, 0, 0, 0}
	img.rgb.Set(x, y, color)
}

func (img *ColorImage) Bounds() image.Rectangle {
	return img.rgb.Bounds()
}
