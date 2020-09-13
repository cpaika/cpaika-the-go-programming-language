// server1 is a minimal echo server, yay
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

// building the variable palette, which is an array of two colors
//var palette = []color.Color{color.White, color.Black}

var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xf0, 0xff, 0x00, 0xff}, color.RGBA{0x0f, 0xbb, 0xf0, 0xff}}

//first and second index of the palette
// I dont love how this works
// interesting way to declare constants tbh I don't really get it
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	fmt.Println("Starting the server on port 8080")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil)) //Why does this take localhost and not a port?
}

func handler(w http.ResponseWriter, r *http.Request) {
	var cycles = 5
	keys, ok := r.URL.Query()["cycles"]

	if ok && len(keys[0]) >= 1 {
		cycles, _ = strconv.Atoi(keys[0])
	}
	lissajous(w, cycles)
}

func lissajous(out io.Writer, cycles int) {
	// you can have const declarations within functions, interesting
	const (
		res     = .001 //angular resolution
		size    = 100  //image canvas covers
		nframes = 256  // number of animation frames
		delay   = 1    //delay between frames in 10 ms units
	)

	// How do computers do random numbers anyway?
	// Sometimes hardware harvesting entropy in nature, sometimes computational randomness
	freq := rand.Float64() * 3.0 // Frequency of the y oscillator

	// Is this creating a new object or a struct?
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // using .0 because thats how you tell it its a float
	for i := 0; i < nframes; i++ {
		// I wonder why plus one here
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		var colorIndex uint8
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			colorIndex = (colorIndex + 1) % 4
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
