package lcd1602

import (
	"time"

	"gitlab.dutchsec.com/pim.van.hespen/go-lcd1602/stringutils"
)

func (lcd *LCD) Slide(s string, line uint8, slideIn, leftSide bool) {
	off := 0
	if slideIn {
		if leftSide {
			off -= lcd.LineWidth
		} else {
			off += lcd.LineWidth
		}
	}
	j := 0
	for i := 0; i <= lcd.LineWidth; i++ {
		j = i
		//slide out to left OR slide in fromright
		if (!slideIn && leftSide) || (slideIn && !leftSide) {
			j *= -1
		}
		sOffset := stringutils.Offset(s, off+j)
		lcd.WriteLine(sOffset, line)
		time.Sleep(SLIDE_SPEED_DELAY)
	}
}

func (lcd *LCD) SlideInLeft(s string, line uint8) {
	lcd.Slide(s, line, true, true)
}

func (lcd *LCD) SlideInRight(s string, line uint8) {
	lcd.Slide(s, line, true, false)
}

func (lcd *LCD) SlideOutLeft(s string, line uint8) {
	lcd.Slide(s, line, false, true)
}

func (lcd *LCD) SlideOutRight(s string, line uint8) {
	lcd.Slide(s, line, false, false)
}
