package lcd1602

import (
	"errors"
	"fmt"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

const (
	RS_DATA        = true  //sending data
	RS_INSTRUCTION = false //sending an instruction

	LINE_1 = uint8(0x80) // address for the 1st line
	LINE_2 = uint8(0xC0) // address for the 2nd line
)

var (
	VERBOSITY = 0

	ENABLE_DELAY = 1 * time.Microsecond

	SLIDE_SPEED_DELAY = 10 * time.Millisecond //lower == faster

	EXECUTION_TIME_DEFAULT     = 38 * time.Microsecond
	EXECUTION_TIME_RETURN_HOME = 1520 * time.Microsecond
)

type LCD struct {
	RS, E     rpio.Pin
	DataPins  []rpio.Pin
	LineWidth int
}

func init() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func Close() {
	rpio.Close()
}

func New(rs, e int, data []int, linewidth int) *LCD {

	datalength := len(data)
	if datalength != 4 && datalength != 8 {
		errors.New("LCD requires four or eight datapins!")
	}

	datapins := make([]rpio.Pin, 0)

	for _, d := range data {
		datapins = append(datapins, rpio.Pin(d))
	}

	lcd := &LCD{
		rpio.Pin(rs),
		rpio.Pin(e),
		datapins,
		linewidth,
	}
	return lcd
}

//Init initiates the LCD
func (l *LCD) Initialize() {

	l.initPins()
	l.Reset()
	l.EntryModeSet(true, false)
	l.DisplayMode(true, false, false) // Display, Cursor, Blink

	l.Write(0x28, RS_INSTRUCTION) // 00101000 - Set DDRAM Address

	l.Clear() // clear screen
}

//ReturnHome function returns the cursor to home
func (l *LCD) ReturnHome() {
	l.Write(0x02, RS_INSTRUCTION)
	time.Sleep(EXECUTION_TIME_RETURN_HOME)
}

//EntryModeSet function
func (l *LCD) EntryModeSet(increment, shift bool) {
	instruction := uint8(0x04)
	if increment {
		instruction |= 0x02
	}
	if shift {
		instruction |= 0x01
	}
	l.Write(instruction, RS_INSTRUCTION)
}

//DisplayMode function set the display modes
func (l *LCD) DisplayMode(display, cursor, blink bool) {
	instruction := uint8(0x08)

	if display {
		instruction |= 0x04
	}
	if cursor {
		instruction |= 0x02
	}
	if blink {
		instruction |= 0x01
	}
	l.Write(instruction, RS_INSTRUCTION)
}

//Clear function clears the screen
func (l *LCD) Clear() {
	l.Write(0x01, RS_INSTRUCTION)
}

func (l *LCD) WriteLines(lines ...string) {
	if len(lines) > 0 {
		l.WriteLine(lines[0], LINE_1)
	}
	if len(lines) > 1 {
		l.WriteLine(lines[1], LINE_2)
	}
}

//WriteLine function writes a single line fo text to the LCD
//if line length exceeds the linelength of the LCD, aslice will be used
func (l *LCD) WriteLine(s string, line uint8) {

	s = fmt.Sprintf("%16s", s)

	l.Write(line, RS_INSTRUCTION)

	for _, c := range s {
		l.Write(uint8(c), RS_DATA)
	}
}

//Write function writes data to the LCD
func (l *LCD) Write(data uint8, mode bool) {
	if mode {
		l.RS.High()
	} else {
		l.RS.Low()
	}

	if len(l.DataPins) == 4 {
		// ofsetfor highest order bits
		base := uint8(0x10)
		for i, dataPin := range l.DataPins {
			setBitToPin(dataPin, data, base<<uint8(i))
		}
		l.enable(EXECUTION_TIME_DEFAULT)
		// lowest order bits
		base = uint8(0x01)
		for i, dataPin := range l.DataPins {
			setBitToPin(dataPin, data, base<<uint8(i))
		}
	} else {
		// all bits
		base := uint8(0x01)
		for i, dataPin := range l.DataPins {
			setBitToPin(dataPin, data, base<<uint8(i))
		}
	}
	l.enable(EXECUTION_TIME_DEFAULT)
}

//setBitToPin function sets given pin to a bit value from a given data int
func setBitToPin(pin rpio.Pin, data, position uint8) {
	if data&position == position {
		pin.High()
	} else {
		pin.Low()
	}
}

//Enable function sets the 'Enable'-pin high, and low to enable 2Xa single write sequence
func (l *LCD) enable(executionTime time.Duration) {
	time.Sleep(ENABLE_DELAY)
	l.E.High()
	time.Sleep(ENABLE_DELAY)
	l.E.Low()
	time.Sleep(executionTime)
}

//private init instruction for LCD
func (l *LCD) Reset() {
	//init sequence
	l.Write(0x33, RS_INSTRUCTION)
	l.Write(0x32, RS_INSTRUCTION)
}

func (l *LCD) initPins() {
	l.RS.Output()
	l.E.Output()
	for _, d := range l.DataPins {
		d.Output()
	}
}
