package gif2lcd

import (
	"image"
	"image/color"
	"image/gif"
	"os"
	"time"

	lcd "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/stringutils"
)

/*
This code is under construction
USE AT YOUR OWN RISK!
*/

func SliceToHex(colors [5]color.Color, treshold int) uint8 {
	result := uint8(0)
	for index, c := range colors {
		r, g, b, _ := c.RGBA()
		lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)

		grayness := uint8(lum / 256) // keep highest 8 bits
		//		fmt.Printf("\t\t%10d %10d %v\n", int64(lum), grayness, grayness > 0x80)
		if grayness > uint8(treshold) {
			result |= (uint8(1) << uint8(index))
		}
		if index >= 4 {
			break
		}
	}
	return result
}

func PxToChar(img *image.Paletted, xBase, yBase, treshold int) lcd.Character {
	result := lcd.Character{}
	for yOffset := 0; yOffset < 8; yOffset++ {
		result[yOffset] = SliceToHex([5]color.Color{
			img.At(xBase+4, yBase+yOffset),
			img.At(xBase+3, yBase+yOffset),
			img.At(xBase+2, yBase+yOffset),
			img.At(xBase+1, yBase+yOffset),
			img.At(xBase+0, yBase+yOffset),
		}, treshold)
	}
	return result
}

func Chrmap(img *image.Paletted, treshold int) []lcd.Character {
	chrmap := []lcd.Character{}

	for y := 0; y < 2; y++ {
		for x := 0; x < 4; x++ {
			chrmap = append(chrmap, PxToChar(img, 5*x, 8*y, treshold))
		}
	}

	return chrmap
}

func BeamToLcd(img *image.Paletted, theLcd *lcd.LCD, delay time.Duration) {

	// TODO remove hard coding
	iterations := 20 //int(delay.Nanoseconds() / (1000 * 1000))

	itertime := delay / time.Duration(iterations)

	offset := 20

	effect := 0.7

	then := time.Now()
	for x := 1; x <= iterations; x++ {

		preoffset := offset + (255-offset)/(x+1)
		treshold := int(float64(preoffset) * effect)

		theLcd.SetCustomCharacters(Chrmap(img, treshold))

		now := time.Now()
		elapsed := now.Sub(then)

		for elapsed > itertime {
			x++
			elapsed -= itertime
		}
		time.Sleep(itertime - elapsed)
	}
}

//ShowGif WARNING: DO NOT USE THIS.
//It was written merely for testing purposes
//It may wear out your LCDs CGRAM
func ShowGif(source string, theLcd *lcd.LCD) {
	input, err := os.Open(source)
	defer input.Close()
	if err != nil {
		panic(err)
	}
	g, err := gif.DecodeAll(input)
	if err != nil {
		panic(err)
	}

	theLcd.WriteLines(
		stringutils.Center("|\x00\x01\x02\x03|", 16),
		stringutils.Center("|\x04\x05\x06\x07|", 16))

	for idx, img := range g.Image {
		sleepy := time.Duration(int64(g.Delay[idx])) * (time.Second / 100)
		BeamToLcd(img, theLcd, sleepy)
	}

}
