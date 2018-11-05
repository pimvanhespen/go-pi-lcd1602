package animations

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randRune() rune {
	return letterRunes[rand.Intn(len(letterRunes))]
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = randRune()
	}
	return string(b)
}

func garble(s string, current, max int) string {
	a := current / (max / len(s))
	r := s[:a] + randStringRunes(len(s)-a)
	return r
}

func garblevert(s string, current, max int) string {
	a := current / (max / len(s))
	r := randStringRunes(len(s)-a) + s[len(s)-a:]
	return r
}

func NewGarbleLeftSimple(s string) Animation {
	// hardcoded default
	// use NewGarbleLeft for extended options
	return NewGarbleLeft(s, 8, 10*time.Millisecond)
}

func NewGarbleLeft(s string, iters int, delay time.Duration) Animation {
	return &GarbleAnimation{
		source: s,
		max:    iters * len(s),
		delay:  delay,
		fn:     garble,
	}
}

func NewGarbleRightSimple(s string) Animation {
	// hardcoded default
	// use NewGarbleRight for extended options
	return NewGarbleRight(s, 8, 10*time.Millisecond)
}

func NewGarbleRight(s string, iters int, delay time.Duration) Animation {
	return &GarbleAnimation{
		source: s,
		max:    iters * len(s),
		delay:  delay,
		fn:     garblevert,
	}
}

type GarbleAnimation struct {
	source  string
	max     int
	delay   time.Duration
	fn      func(string, int, int) string
	current int
}

func (g *GarbleAnimation) Content() string {
	g.current++
	return g.fn(g.source, g.current, g.max)
}
func (g *GarbleAnimation) Done() bool {
	return g.current >= g.max
}
func (g *GarbleAnimation) Delay() {
	time.Sleep(g.delay)
}
