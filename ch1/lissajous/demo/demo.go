package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"strconv"
)

const (
	size = 128
)

func main() {
	p := 3.0
	q := 4.0
	phi := 0.0

	if len(os.Args) >= 2 {
		p, _ = strconv.ParseFloat(os.Args[1], 64)
	}
	if len(os.Args) >= 3 {
		q, _ = strconv.ParseFloat(os.Args[2], 64)
	}
	if len(os.Args) >= 4 {
		phi, _ = strconv.ParseFloat(os.Args[3], 64)
	}

	palett := []color.Color{color.RGBA{0xee, 0xee, 0xee, 0xff}, color.RGBA{0xff, 0, 0, 0xff}}
	rec := image.Rect(0, 0, 2*size, 2*size)
	img := image.NewPaletted(rec, palett)

	s := float64(size - 10) // 比例缩放
	for t := 0.0; t <= 2*math.Pi; t += 0.0001 {
		x := math.Sin(p * t)
		y := math.Sin(q*t + phi)
		img.SetColorIndex(size+int(s*x+0.5), size+int(s*y+0.5), 1)
	}

	png.Encode(os.Stdout, img)
}
