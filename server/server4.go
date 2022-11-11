package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

var (
	cycles     = 5     // number of complete x oscillator revolutions
	res        = 0.001 // angular resolution
	size       = 100   // image canvas covers [-size..+size]
	nframes    = 64    // number of animation frames
	delay      = 8     // delay between frames in 10ms units
	rand_color = false // random color
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		// reset variables
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
		rand_color = false

		for k, v := range r.Form {
			if k == "cycles" {
				num, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Println(err)
				}
				cycles = num
			}

			if k == "res" {
				num, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Println(err)
				}
				res = float64(num)
			}

			if k == "size" {
				num, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Println(err)
				}
				size = num
			}

			if k == "nframes" {
				num, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Println(err)
				}
				nframes = num
			}

			if k == "delay" {
				num, err := strconv.Atoi(v[0])
				if err != nil {
					fmt.Println(err)
				}
				delay = num
			}

			if k == "rand_color" {
				num, err := strconv.ParseBool(v[0])
				if err != nil {
					fmt.Println(err)
				}
				rand_color = num
			}
		}

		lissajous(w)
	}) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func lissajous(out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	palette := []color.Color{color.White, color.Black}
	for i := 0; i < nframes; i++ {
		if rand_color {
			// 随机颜色
			var randUint16 uint16 = uint16(math.Floor(rand.Float64() * 0xffffff))
			var randUint16_2 uint16 = uint16(math.Floor(rand.Float64() * 0xffffff))
			// fmt.Printf("输出: %x\n", randUint16)

			palette = []color.Color{color.Gray16{randUint16}, color.Gray16{randUint16_2}}
		}

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
