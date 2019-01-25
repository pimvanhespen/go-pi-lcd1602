package animations

import "fmt"

type NoAnimation struct {
	source string
	done   bool
}

func (n *NoAnimation) Width(width int) {
	n.source = fmt.Sprintf(fmt.Sprintf("%%%ds", width), n.source)
}
func (n *NoAnimation) Done() bool { return n.done }
func (n *NoAnimation) Delay()     {}
func (n *NoAnimation) Content() string {
	n.done = true
	return n.source
}

func None(s string) Animation {
	return &NoAnimation{
		source: s,
	}
}
