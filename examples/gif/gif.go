package main

import (
	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/gif2lcd"
	synchronizedLcd "github.com/pimvanhespen/go-pi-lcd1602/synchronized"
)

func main() {
	// !! WARNING! USE AT OWN RISK !!
	// !! FLASHING THE CGRAM MIGHT NOT BE GOOD FOR YOUR LCD !!
	lcd := lcd1602.New(
		7,                     //rs
		8,                     //enable
		[]int{25, 24, 23, 18}, //datapins
		16,                    //lineSize
	)
	lcd.Initialize()

	lcdi := synchronizedLcd.NewSynchronizedLCD(lcd)
	lcdi.WriteLines("Go Rpi LCD 1602", "git/PimvanHespen")

	gif2lcd.ShowGif("test.gif", lcdi)
	lcd1602.Close()
}
