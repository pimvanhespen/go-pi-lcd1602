package animations

import (
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

func NewSlideInLeft(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: -len(s),
		max:     0,
		fn:      slideInLeft,
		delay:   time.Millisecond * 35,
	}
}

func NewSlideInLeftX(s string, delay time.Duration) Animation {
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

func NewSlideInRight(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideInRight,
		delay:   time.Millisecond * 35,
	}
}

func slideInRight(s string, current int) string {
	offset := len(s) - current
	return stringutils.Offset(s, offset)
}

func NewSlideOutLeft(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideOutLeft,
		delay:   time.Millisecond * 35,
	}
}

func slideOutLeft(s string, current int) string {
	offset := -current
	return stringutils.Offset(s, offset)
}

func NewSlideOutRight(s string) Animation {
	return &SlideAnimation{
		source:  s,
		current: 0,
		max:     len(s),
		fn:      slideOutRight,
		delay:   time.Millisecond * 35,
	}
}

func NewSlideOutRightX(s string, delay time.Duration) Animation {
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
