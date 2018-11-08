package main

import (
	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/animations"
	"github.com/pimvanhespen/go-pi-lcd1602/stringutils"
)

func main() {
	lcd := lcd1602.New(7, 8, []int{25, 24, 23, 18}, 16)
	lcd.Initialize()

	animations := []animations.Animation{
		animations.NewNoAnimation(stringutils.Center("no animation", 16)),
		animations.NewGarbleLeftSimple(stringutils.Center("garble left", 16)),
		animations.NewGarbleRightSimple(stringutils.Center("garble right", 16)),
		animations.NewSlideInLeft(stringutils.Center("slide in left", 16)),
		animations.NewSlideInRight(stringutils.Center("slide in right", 16)),
		animations.NewSlideOutLeft(stringutils.Center("slide out left", 16)),
		animations.NewSlideOutRight(stringutils.Center("slide out right", 16)),
		animations.NewNoAnimation(""),
		animations.NewGarbleLeftSimple("   github.com   "),
		animations.NewGarbleRightSimple(" /pimvanhespen  "),
	}

	for index, animation := range animations {
		line := lcd1602.LINE_2
		if index%2 == 0 {
			line = lcd1602.LINE_1
		}
		//<-lcd.Animate(animation, line) //shorter version of next 2 lines
		wait := lcd.Animate(animation, line)
		<-wait
	}

	lcd1602.Close()
}
