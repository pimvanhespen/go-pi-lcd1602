package main

import (
	"time"

	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
)

func main() {
	lcd := lcd1602.New(
		7,                     //rs
		8,                     //enable
		[]int{25, 24, 23, 18}, //datapins
		16,                    //lineSize
	)
	lcd.Initialize()

	lines := []string{
		"Go Rpi LCD 1602",
		"git/PimvanHespen",
		"Hello, World!",
	}

	lcd.WriteLines(lines[0])

	top, bottom := lines[2], lines[0]

	for x := 0; x < (3 * len(lines)); x++ {
		lcd.SlideOutLeft(top, lcd1602.LINE_1)
		lcd.SlideOutRight(bottom, lcd1602.LINE_2)

		top, bottom = lines[x%len(lines)], lines[(x+1)%len(lines)]
		lcd.SlideInRight(top, lcd1602.LINE_1)
		lcd.SlideInLeft(bottom, lcd1602.LINE_2)

		time.Sleep(1 * time.Second)
	}

	lcd1602.Close()
}
