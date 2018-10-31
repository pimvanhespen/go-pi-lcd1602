package stringutils

import (
	"fmt"
	"math"
	"strings"
)

func Center(s string, width int) string {
	rem := width - len(s)
	if rem > 1 {
		leftF := fmt.Sprintf("%%-%ds", width)              // format left
		rightF := fmt.Sprintf("%%%ds", len(s)+(rem-rem/2)) // format right
		right := fmt.Sprintf(rightF, s)                    // right half
		s = fmt.Sprintf(leftF, right)                      // final string
	}
	return s
}

func Offset(s string, offset int) string {
	// no offset, string wont change
	if offset == 0 {
		return s
	}
	strLen := len(s)
	//offset greater than length of string, return an emtpry string
	if offset >= strLen || offset <= (-1*strLen) {
		return strings.Repeat(" ", strLen)
	}

	if offset < 0 {
		//offset negative, string goes left

		absOffset := int(math.Abs(float64(offset)))
		return s[absOffset:] + strings.Repeat(" ", absOffset)
	} else {
		//offset positive, string goes right
		return strings.Repeat(" ", offset) + s[:strLen-offset]
	}
}
