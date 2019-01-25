package main

import (
	"time"

	lcd1602 "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/synchronized"
)

func main() {
	lcdi := lcd1602.New(
		10,                   //rs
		9,                    //enable
		[]int{6, 13, 19, 26}, //datapins
		16,                   //lineSize
	)
	lcd := synchronized.NewSynchronizedLCD(lcdi)
	lcd.Initialize()
	lcd.WriteLines("Go Rpi LCD 1602", "git/PimvanHespen")
	time.Sleep(1 * time.Second)
	lcd.Clear()
	lcd.Close()
}
