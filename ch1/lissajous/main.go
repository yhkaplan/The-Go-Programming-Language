// Generates GIF animations of random Lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// A composite literal, a compact notation for instantiating any of
// Go's composite types from a sequence of element values
// This is a slice
var palette = []color.Color{color.Black, color.RGBA{0, 255, 0, 1}, color.RGBA{255, 0, 0, 1}, color.RGBA{0, 0, 255, 1}}

const (
	blackIndex = 0 // First color in palette
	greenIndex = 1 // Second color in palette
	redIndex   = 2
	blueIndex  = 3
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // # of complete x oscillations
		res     = 0.001 // angular resolution
		size    = 100   // image canvas
		nframes = 64    // Animation frames
		delay   = 8     //delay between frame in 10 ms units
	)

	freq := rand.Float64() * 3.0 // relative freq of oscillator
	// This is a struct
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// This outer loop runs for 64 iterations, each producing a single
	// frame of animation. It creates a new 201x201 image with a palette
	// of 2 colors: black and white
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// Each pass through this inner loop generates a new image by setting some pixels
		// to be black (all the others are set to white previously)
		// The result is appended to a list of frames, along with a delay of 80 ms
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			randomColor := uint8(rand.Intn(4))

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), randomColor)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Finally, the sequence of frames and delays is encoded into GIF and written to the
	// output stream 'out'
	gif.EncodeAll(out, &anim) // Ignoring encoding errors
}
