package main

import (
	"flag"
	"github.com/chr1sj0nes/go-fractal"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	width := flag.Int("w", 600, "Output width / pixels")
	height := flag.Int("h", 400, "Output height / pixels")
	format := flag.String("f", "png", "Output image format.")
	output := flag.String("o", "-", "Outfile file.")

	flag.Parse()

	var encode func(w io.Writer, m image.Image) error

	switch *format {
	case "png":
		encode = png.Encode
	case "gif":
		encode = func(w io.Writer, m image.Image) error {
			return gif.Encode(w, m, nil)
		}
	case "jpg", "jpeg":
		encode = func(w io.Writer, m image.Image) error {
			return jpeg.Encode(w, m, nil)
		}
	default:
		log.Fatalf("Unsupported format: %s", *format)
	}

	w := os.Stdout
	if *output != "-" {
		f, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}

		defer f.Close()
		w = f
	}

	r := image.Rect(0, 0, *width, *height)
	img := image.NewRGBA(r)

	min := complex(-2, -1)
	max := complex(1, 1)

	fract.Mandelbrot(img, fract.ColorBinary, min, max)

	if err := encode(w, img); err != nil {
		log.Fatal(err)
	}
}
