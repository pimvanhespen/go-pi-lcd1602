package synchronized

import (
	"sync"

	lcd "github.com/pimvanhespen/go-pi-lcd1602"
	"github.com/pimvanhespen/go-pi-lcd1602/animations"
)

type SynchronizedLCD struct {
	lcd.LCDI
	line1, line2 sync.Mutex
}

func NewSynchronizedLCD(l lcd.LCDI) *SynchronizedLCD {
	l.Initialize()
	return &SynchronizedLCD{
		l, sync.Mutex{}, sync.Mutex{},
	}
}
func (l *SynchronizedLCD) WriteLines(lines ...string) {
	if len(lines) > 0 {
		l.line1.Lock()
		l.WriteLine(lines[0], lcd.LINE_1)
		l.line1.Unlock()
	}
	if len(lines) > 1 {
		l.line2.Lock()
		l.WriteLine(lines[1], lcd.LINE_2)
		l.line2.Unlock()
	}
}

func (l *SynchronizedLCD) Animate(animation animations.Animation, line lcd.LineNumber) chan bool {
	done := make(chan bool, 1)
	var mut sync.Mutex
	//TODO: fix hardcoding (catch error for unrecognizedlines, etc..)
	if line == lcd.LINE_1 {
		mut = l.line1
	} else {
		mut = l.line2
	}

	mut.Lock()
	go func() {
		animation.Width(l.LCDI.Width())
		for !animation.Done() {
			s := animation.Content()
			l.WriteLine(s, line)
			animation.Delay()

		}
		mut.Unlock()
		done <- true
	}()

	return done
}
