package lcd1602

import (
	"errors"
	"fmt"
	"os"
	"sync"
	"time"

	rpio "github.com/stianeikeland/go-rpio"
)

const (
	RS_DATA        = true  //sending data
	RS_INSTRUCTION = false //sending an instruction

	LINE_1 = LineNumber(0x80) // address for the 1st line
	LINE_2 = LineNumber(0xC0) // address for the 2nd line
)

var (
	ENABLE_DELAY               = 1 * time.Microsecond
	EXECUTION_TIME_DEFAULT     = 40 * time.Microsecond
	EXECUTION_TIME_RETURN_HOME = 1520 * time.Microsecond
)

var lines = map[int]LineNumber{
	1: LINE_1,
	2: LINE_2,
}

//global used to ensure the rpio library is nitialized befure using it.
var rpioPrepared = false

type LineNumber uint8

type Character [8]uint8

type LCD struct {
	RS, E               rpio.Pin
	DataPins            []rpio.Pin
	LineWidth           int
	writelock, linelock sync.Mutex
}

type LCDI interface {
	Initialize()
	ReturnHome()
	EntryModeSet(bool, bool)
	DisplayMode(bool, bool, bool)
	Clear()
	Reset()
	Write(uint8, bool)
	WriteLine(string, LineNumber)
	CreateChar(uint8, Character)
	Width() int
	Close()
}

func SetCustomCharacters(l LCDI, characters []Character) {
	for index, chr := range characters {
		offset := 8 - len(characters) + index
		if offset < 0 {
			continue
		}
		l.CreateChar(uint8(offset), chr)
	}
}

//function should be called before executing any other code!
func Open() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rpioPrepared = true
}

func Close() {
	if rpioPrepared {
		rpio.Close()
	}
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

	l := &LCD{
		RS:        rpio.Pin(rs),
		E:         rpio.Pin(e),
		DataPins:  datapins,
		LineWidth: linewidth,
	}
	l.initPins()
	return l
}

func (l *LCD) Close() {}
func (l *LCD) Width() int {
	return l.LineWidth
}

//Init initiates the LCD
func (l *LCD) Initialize() {
	l.Reset()

	l.EntryModeSet(true, false)
	l.DisplayMode(true, false, false) // Display, Cursor, Blink

	l.Write(0x28, RS_INSTRUCTION) // 00101000 - Set DDRAM Address
	l.ReturnHome()

	l.Clear() // clear screen
	//init time...
	time.Sleep(10 * time.Millisecond)
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

//WriteLine function writes a single line fo text to the LCD
//if line length exceeds the linelength of the LCD, aslice will be used
func (l *LCD) WriteLine(s string, line LineNumber) {
	l.linelock.Lock()
	defer l.linelock.Unlock()
	frmt := fmt.Sprintf("%%%ds", l.LineWidth)
	s = fmt.Sprintf(frmt, s)

	s = s[:l.LineWidth]

	l.Write(uint8(line), RS_INSTRUCTION)

	for _, c := range s {
		l.Write(uint8(c), RS_DATA)
	}
}

//Write function writes data to the LCD
func (l *LCD) Write(data uint8, mode bool) {
	l.writelock.Lock()
	defer l.writelock.Unlock()

	if mode {
		l.RS.High()
	} else {
		l.RS.Low()
	}

	for _, p := range l.DataPins {
		p.Low()
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

func (l *LCD) CreateChar(position uint8, data Character) {
	if position > 7 {
		//error
		return
	}
	l.Write(0x40|(position<<3), false)
	for _, x := range data {
		l.Write(x, true)
	}
}

//Reset resets the lcd
func (l *LCD) Reset() {
	//init sequence
	l.Write(0x33, RS_INSTRUCTION)
	time.Sleep(EXECUTION_TIME_DEFAULT)
	l.Write(0x32, RS_INSTRUCTION)
	time.Sleep(EXECUTION_TIME_DEFAULT)
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

func (l *LCD) initPins() {
	if !rpioPrepared {
		Open()
	}
	l.RS.Output()
	l.E.Output()
	for _, d := range l.DataPins {
		d.Output()
	}
}
