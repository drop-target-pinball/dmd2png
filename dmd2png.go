package main

import (
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	data, err := ioutil.ReadFile("p-roc.dmd")
	if err != nil {
		panic(err)
	}

	palette := make([]color.Color, 0)
	for i := 0; i <= 0xff; i += 0x11 {
		c := uint8(i)
		palette = append(palette, color.RGBA{c, c, c, 0xff})
	}
	img := image.NewPaletted(image.Rect(0, 0, 128, 32), palette)

	// skip header
	data = data[4+4+4+4:]
	x := 0
	y := 0
	for _, d := range data {
		d = d & 0xf
		img.SetColorIndex(x, y, uint8(d))
		x += 1
		if x >= 128 {
			x = 0
			y += 1
		}
	}

	f, err := os.Create("p-roc.png")
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
