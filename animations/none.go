package animations

type NoAnimation struct {
	source string
	done   bool
}

func (n *NoAnimation) Done() bool { return n.done }
func (n *NoAnimation) Delay()     {}
func (n *NoAnimation) Content() string {
	n.done = true
	return n.source
}

func NewNoAnimation(s string) Animation {
	return &NoAnimation{
		source: s,
	}
}
