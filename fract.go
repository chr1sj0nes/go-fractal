// Functions for generating images of fractal sets
package fract

import (
	"image"
	"image/color"
	"image/draw"
	"math/cmplx"
)

// Iterations after which we are "converged"
const MaxIterations = 10000

// Color the image using iterations to divergence
type Colorize func(iterations int) color.Color

// Generate Mandelbrot set
func Mandelbrot(img draw.Image, col Colorize, min, max complex128) {
	b := img.Bounds()
	dr := real(max-min) / float64(b.Dx())
	di := imag(max-min) / float64(b.Dy())

	ch := make(chan bool, b.Dy())

	for y := 0; y < b.Dy(); y++ {
		go func(y int) {
			for x := 0; x < b.Dx(); x++ {
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
		}(y)
	}

	// wait for all go routines to finish
	for i := 0; i < b.Dy(); i++ {
		<-ch
	}
}

func ColorBinary(iterations int) color.Color {
	if iterations == MaxIterations {
		return color.Black
	}

	return color.White
}

func CountBlack(img image.Image) uint {
	b := img.Bounds()

	r0, g0, b0, a0 := color.Black.RGBA()

	count := uint(0)
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			if (r == r0) && (g == g0) && (b == b0) && (a == a0) {
				count++
			}
		}
	}

	return count
}
