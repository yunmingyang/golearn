package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)


func main(){
	const (
		xmin, xmax, ymin, ymax = -2, +2, -2, +2
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0,0, width, height))

	for py := 0; py < height; py++ {
		y := float64(py) / height * (ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px) / width * (xmax - xmin) + xmin
			z := complex(x, y)
			// Image point represents complex value z
			img.Set(px, py, mandelbro(z))
		}
	}

	// output to stdout
	png.Encode(os.Stdout, img)
	// f, err := os.Create("test.png")
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "file creation failed as: %s", err)
	// 	os.Exit(1)
	// }

	// png.Encode(f, img)
	// f.Close()
}

func mandelbro(z complex128) color.Color{
	const (
		iterations = 200
		contrast = 15
	)

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}

	return color.Black
}