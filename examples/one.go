package main

import (
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
	lcd.WriteLines("Go Rpi LCD 1602", "git/PimvanHespen")
	lcd1602.Close()
}
