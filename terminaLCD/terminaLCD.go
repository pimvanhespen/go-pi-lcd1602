package terminaLCD

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/fatih/color"
	lcd "github.com/pimvanhespen/go-pi-lcd1602"
)

type TerminalLCD struct {
	//path string
	file         *os.File
	linewidth    int
	line1, line2 string
	lock1, lock2 sync.Mutex
}

func (f *TerminalLCD) Initialize() {

	//TODO Pim van Hespen: move this to config
	f.linewidth = 16

	dir, err1 := os.Getwd()
	if err1 != nil {
		panic(err1)
	}

	path := dir + "/LCD"

	if _, err := os.Stat(path); os.IsNotExist(err) {

		_, err := os.Create(path)
		if err != nil {
			panic(err)
		}
	}

	file, err := os.OpenFile(path, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	//f.path = path

	//display command for viewing output to user
	fmt.Println("The Terminal LCD is visible with the following command on Linux")
	fmt.Printf("\n\ttail -f %s\n\n", path)
	f.file = file

	f.file.Seek(0, 0)
	f.file.Truncate(0)
	f.file.Sync()

	// overwrite default output for goterm
	//	tm.Output = bufio.NewWriter(f.file)
}
func (f *TerminalLCD) Clear() {
	f.WriteLine("", lcd.LINE_1)
	f.WriteLine("", lcd.LINE_2)
}
func (f *TerminalLCD) EntryModeSet(a, b bool)   {}
func (f *TerminalLCD) DisplayMode(a, b, c bool) {}
func (f *TerminalLCD) Reset()                   {}
func (f *TerminalLCD) Width() int {
	return 16
}
func (f *TerminalLCD) Write(cmd uint8, mode bool)               {}
func (f *TerminalLCD) CreateChar(pos uint8, char lcd.Character) {}
func (f *TerminalLCD) ReturnHome()                              {}
func (f *TerminalLCD) Close() {
	//	f.file.Close()
}

func ReplaceCustomCharacters(s string) string {
	conversion := map[string]string{
		"\x00": "\u2080",
		"\x01": "\u2081",
		"\x02": "\u2082",
		"\x03": "\u2083",
		"\x04": "\u2084",
		"\x05": "\u2085",
		"\x06": "\u2086",
		"\x07": "\u2087",
		" ":    "\u2591",
	}
	for key, value := range conversion {
		s = strings.Replace(s, key, value, -1)
	}
	return s
}

func (f *TerminalLCD) Update() {

	frmt := fmt.Sprintf("%%%ds", f.linewidth)

	//content
	lcdLineOne := fmt.Sprintf(frmt, ReplaceCustomCharacters(f.line1))
	lcdLineTwo := fmt.Sprintf(frmt, ReplaceCustomCharacters(f.line2))

	//unicode points
	ucTop, ucLeft, usRight, ucBottom := "\u2581", "\u2588", "\u2588", "\u2594"
	top := strings.Repeat(ucTop, 18)
	bottom := strings.Repeat(ucBottom, 18)

	//colors
	boldwhiteblack := color.New(color.FgHiWhite, color.BgBlack, color.Bold)
	whitegreen := color.New(color.FgHiWhite, color.BgGreen)
	blackgreen := color.New(color.FgBlack, color.BgGreen)
	yellowgreen := color.New(color.FgYellow, color.BgGreen)
	whiteblue := color.New(color.BgBlue, color.FgHiWhite)

	//apply color to lcd lines
	lcdLineOne = whiteblue.Sprint(lcdLineOne)
	lcdLineTwo = whiteblue.Sprint(lcdLineTwo)

	//string pieces
	emptyPre := boldwhiteblack.Sprint(strings.Repeat(" ", 7))
	lineTrailing := boldwhiteblack.Sprint(strings.Repeat(" ", 4))

	//lines
	preHeadLine := emptyPre + blackgreen.Sprint("\u2981") + yellowgreen.Sprintf(" %s ", strings.Repeat("\u2596", 16)) + blackgreen.Sprint("\u2981") + lineTrailing
	headline := emptyPre + whitegreen.Sprintf(" %s ", top) + lineTrailing
	secondLine := boldwhiteblack.Sprint(" DEBUG ") + whitegreen.Sprintf(" %s", ucLeft) + lcdLineOne + whitegreen.Sprintf("%s ", usRight) + lineTrailing
	thirdLine := boldwhiteblack.Sprint("  LCD\u2122 ") + whitegreen.Sprintf(" %s", ucLeft) + lcdLineTwo + whitegreen.Sprintf("%s ", usRight) + lineTrailing
	bottomLine := emptyPre + blackgreen.Sprint("\u2981") + whitegreen.Sprintf("%s", bottom) + blackgreen.Sprint("\u2981") + lineTrailing

	//spacing
	marginLine := strings.Repeat(" ", 31) //ugly... I Know

	result := "\033[2J\n" + strings.Join([]string{
		boldwhiteblack.Sprintf("%s", marginLine),
		boldwhiteblack.Sprintf("%s", marginLine),
		preHeadLine,
		headline,
		secondLine,
		thirdLine,
		bottomLine,
		boldwhiteblack.Sprintf("%s", marginLine),
		boldwhiteblack.Sprintf("%s", marginLine),
	}, "\n")

	ln, err := f.file.Stat()
	if err != nil {
		panic(err)
	}

	// ensure file does not get too large....
	if ln.Size() > 10000 {
		f.file.Seek(0, 0)
		f.file.Truncate(0)
	}
	f.file.Write([]byte(result))

	f.file.Sync()

}

func (f *TerminalLCD) WriteLine(s string, line lcd.LineNumber) {
	if line == lcd.LINE_1 {
		f.line1 = s
	} else {
		f.line2 = s
	}
	f.Update()
}
