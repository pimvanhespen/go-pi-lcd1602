package main

import (
	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/gif2lcd"
)

func main() {
	lcd := lcd1602.New(
		7,                     //rs
		8,                     //enable
		[]int{25, 24, 23, 18}, //datapins
		16,                    //lineSize
	)
	lcd.Initialize()
	lcd.WriteLines("Go Rpi LCD 1602", "git/PimvanHespen")
	gif2lcd.ShowGif("test.gif", lcd)
	lcd1602.Close()
}
