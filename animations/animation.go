package animations

type Animation interface {
	Width(int)
	Content() string
	Delay()
	Done() bool
}
