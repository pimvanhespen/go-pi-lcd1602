# Golang  LCD1602 (LCD16x2) for RaspberryPi 
[![GoDoc](https://godoc.org/github.com/pimvanhespen/go-pi-lcd1602?status.svg)](https://godoc.org/github.com/pimvanhespen/go-pi-lcd1602)
[![Go Report Card](https://goreportcard.com/badge/github.com/pimvanhespen/go-pi-lcd1602)](https://goreportcard.com/report/github.com/pimvanhespen/go-pi-lcd1602)
[![license](https://img.shields.io/github/license/pimvanhespen/go-pi-lcd1602.svg)](https://github.com/pimvanhespen/go-pi-lcd1602/blob/master/LICENSE.md)


Use LCD screens in your Go RPi applications!


## Fast
The timing in this library is optimized to run as smoot as possible.
(It takes **~40 microseconds** to write one character to the LCD, opposed to many online examples taking **5-10 milliseconds**).

## Animated
You can use **Animations** (see animations, and examples/animations.go) to slide text into and out of the LCD.
You can also create your own animations by implementing the `Animation` interface.

## Virtual LCD
I wrote a virtual representation of an LCD screen. You can use it for debugging on de command line, or when your LCD is broken... 

## Usage
### Import

```go
import "github.com/pimvanhespen/go-pi-lcd1602"
```
Also checkout the [examples](https://github.com/pimvanhespen/go-pi-lcd1602/tree/master/examples)!

### Short example
```go
    // write code to your LCD in 5 simple steps!
    // 1. Define the LCD config (which pins are used)
	lcdi := lcd1602.New(
		10,                   //rs
		9,                    //enable
		[]int{6, 13, 19, 26}, //datapins
		16,                   //lineSize
	)
    // 2. Create a synchronized LCD (for writing both lines easily)
	lcd := synchronized.NewSynchronizedLCD(lcdi)
    // 3. Init the LCD
	lcd.Initialize()
    // 4. Write text to the LCD
	lcd.WriteLines("Go Rpi LCD 1602", "git/PimvanHespen")
    // 5. Close the lcd (i.e. clean up GPIO memory)
    lcd.Close()
}
```
## Todo
- lanning on decoupling the LCD from the RaspberryPi GPIO. Allows for users to write/use their own IO wrappers for different hardware solutions

## Special thanks
This library is built using [Stian Eikelands go-rpio library](https://github.com/stianeikeland/go-rpio).

## Changelog
### 25/01/2019
- Decoupled the lcd from synchronization
- Added a VIRTUAL LCD
- rewrote examples to match new code

### 5/11/2018
- Fixed boot bug: The LCD used to randomly show garbled data. It's fixed
- Better implementation of animations. Animations are now locked to a single line, allowing for animating both lines concurrently
- Added an example for the new animations
