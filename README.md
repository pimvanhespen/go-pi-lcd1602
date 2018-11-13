# Golang  LCD1602 (LCD16x2) for RaspberryPi 
[![GoDoc](https://godoc.org/github.com/pimvanhespen/go-pi-lcd1602?status.svg)](https://godoc.org/github.com/pimvanhespen/go-pi-lcd1602)
[![Go Report Card](https://goreportcard.com/badge/github.com/pimvanhespen/go-pi-lcd1602)](https://goreportcard.com/report/github.com/pimvanhespen/go-pi-lcd1602)
[![license](https://img.shields.io/github/license/pimvanhespen/go-pi-lcd1602.svg)](https://github.com/pimvanhespen/go-pi-lcd1602/blob/master/LICENSE)
Use LCD screens in your Go RPi applications!


## Fast
The timing in this library is optimized to run as smoot as possible.
(It takes **~40 microseconds** to write one character to the LCD, opposed to many online examples taking **5-10 milliseconds**).

## Animated
You can use **Animations** (see animations, and examples/animations.go) to slide text into and out of the LCD.
You can also create your own animations by implementing the `Animation` interface.

## Changelog
### 5 / 11 / 2018
- Fixed boot bug: The LCD used to randomly show garbled data. It's fixed
- Better implementation of animations. Animations are now locked to a single line, allowing for animating both lines concurrently
- Added an example for the new animations


## Usage
### Import

```go
import "github.com/pimvanhespen/go-pi-lcd1602"
```
Also checkout the [examples](https://github.com/pimvanhespen/go-pi-lcd1602/tree/master/examples)!

### Short example
```go
func main() {
	//Write text to your LCD in 4 lines of code
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
```
## Special thanks
This library is built using [Stian Eikelands go-rpio library](https://github.com/stianeikeland/go-rpio).
