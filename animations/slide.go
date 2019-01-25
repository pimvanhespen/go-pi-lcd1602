package animations

import (
	"fmt"
	"time"

	"github.com/pimvanhespen/go-pi-lcd1602/stringutils"
)

type SlideAnimation struct {
	source  string
	current int
	delay   time.Duration
	fn      func(string, int) string
	max     int
}

func (s *SlideAnimation) Width(width int) {
	old := s.source
	s.source = fmt.Sprintf(fmt.Sprintf("%%%ds", width), old)
	s.current = (s.current / len(old)) * width
	s.max = (s.max / len(old)) * width
}
func (s *SlideAnimation) Content() string {
	s.current++
	ret := s.fn(s.source, s.current)
	return ret
}
func (s *SlideAnimation) Done() bool {
	return s.current >= s.max
}
func (s *SlideAnimation) Delay() {
	time.Sleep(s.delay)
}

func SlideInLeft(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: -len(s),
		max:     0,
		fn:      slideInLeft,
		delay:   time.Millisecond * 20,
	}
}

func SlideInLeftX(s string, delay time.Duration) Animation {
	return &SlideAnimation{
		source:  s,
		current: -len(s),
		max:     0,
		fn:      slideInLeft,
		delay:   delay,
	}
}

func slideInLeft(s string, current int) string {
	return stringutils.Offset(s, current)
}

func SlideInRight(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideInRight,
		delay:   time.Millisecond * 20,
	}
}

func slideInRight(s string, current int) string {
	offset := len(s) - current
	return stringutils.Offset(s, offset)
}

func SlideOutLeft(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideOutLeft,
		delay:   time.Millisecond * 20,
	}
}

func slideOutLeft(s string, current int) string {
	offset := -current
	return stringutils.Offset(s, offset)
}

func SlideOutRight(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideOutRight,
		delay:   time.Millisecond * 20,
	}
}

func SlideOutRightX(s string, delay time.Duration) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideOutRight,
		delay:   delay,
	}
}

func slideOutRight(s string, current int) string {
	return stringutils.Offset(s, current)
}
