// Functions for generating images of fractal sets
package fract

import (
	"image"
	"image/color"
	"image/draw"
	"math/cmplx"
)

const MaxIterations = 10000

// Color the image using iterations to diverge
type Colorize func(iterations int) color.Color

// Generate Mandelbrot set
func Mandelbrot(img draw.Image, col Colorize, min, max complex128) {
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

				img.Set(x+b.Min.X, y+b.Min.Y, col(n))
			}

			ch <- true
		}(x)
	}

	// wait for all go routines to finish
	for i := 0; i < b.Dx(); i++ {
		<-ch
	}
}

func ColorBinary(iterations int) color.Color {
	if iterations == MaxIterations {
		return color.Black
	} else {
		return color.White
	}
}

func CountBlack(img image.Image) uint {
	b := img.Bounds()
	
	r0, g0, b0, a0 := color.Black.RGBA()

	count := uint(0)
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			
			if (r == r0) && (g == g0) && (b == b0) && (a == a0) {
				count++
			}
		}
	}
	
	return count
}
