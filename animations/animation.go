package animations

type Animation interface {
	Content() string
	Delay()
	Done() bool
}
