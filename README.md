# Golang  LCD1602 (LCD16x2) for RaspberryPi [![GoDoc](https://godoc.org/github.com/pimvanhespen/go-pi-lcd1602?status.svg)](https://godoc.org/github.com/pimvanhespen/lcd1602)
Use LCD screens in your Go RPi applications!


## Fast
The timing in this library is optimized to run as smoot as possible.
(It takes **~40 microseconds** to write one character to the LCD, opposed to many online examples taking **5-10 milliseconds**).

## Fun
You can use **Animations** (see animations.go, and examples/slide.go) to slide text into and out of the LCD.

This library is built using [Stian Eikelands go-rpio library](https://github.com/stianeikeland/go-rpio).





## Usage
### Import

```go
import "github.com/pimvanhespen/go-pi-lcd1602"
```
Also checkout the [examples](https://github.com/pimvanhespen/go-pi-lcd1602/tree/master/examples)!

### Short example
```go
func main() {
	//Outout in your display in 4 lines!
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



