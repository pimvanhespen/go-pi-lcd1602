package main

import (
	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/animations"
	"github.com/pimvanhespen/go-pi-lcd1602/stringutils"
	"github.com/pimvanhespen/go-pi-lcd1602/synchronized"
)

func main() {
	lcd := lcd1602.New(
		10,                   //rs
		9,                    //enable
		[]int{6, 13, 19, 26}, //datapins
		16,                   //lineSize
	)

	lcd.Initialize()

	lcdi := synchronized.NewSynchronizedLCD(lcd)

	animations := []animations.Animation{
		animations.None(stringutils.Center("no animation", 16)),
		animations.GarbleLeftSimple(stringutils.Center("garble left", 16)),
		animations.GarbleRightSimple(stringutils.Center("garble right", 16)),
		animations.SlideInLeft(stringutils.Center("slide in left", 16)),
		animations.SlideInRight(stringutils.Center("slide in right", 16)),
		animations.SlideOutLeft(stringutils.Center("slide out left", 16)),
		animations.SlideOutRight(stringutils.Center("slide out right", 16)),
		animations.GarbleLeftSimple("   github.com   "),
		animations.GarbleRightSimple(" /pimvanhespen  "),
	}

	for index, animation := range animations {
		line := lcd1602.LINE_2
		if index%2 == 0 {
			line = lcd1602.LINE_1
		}
		//<-lcd.Animate(animation, line) //shorter version of next 2 lines
		wait := lcdi.Animate(animation, line)
		<-wait
	}

	lcd1602.Close()
}
