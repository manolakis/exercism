package clock

import (
	"fmt"
	"math"
)

type Clock struct {
	hours   int
	minutes int
}

func New(h, m int) Clock {
	m, j := adjustMinutes(m)

	return Clock{hours: adjustHours(h + j), minutes: m}
}

func adjustHours(h int) int {
	h = h % 24

	if h < 0 {
		return 24 + h
	}

	return h
}

func adjustMinutes(m int) (int, int) {
	var h int

	if math.Abs(float64(m)) > 0 {
		h = m / 60
		m = m % 60
	}

	if m < 0 {
		return 60 + m, h - 1
	}

	return m, h
}

func (c Clock) Add(m int) Clock {
	return New(c.hours, c.minutes+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.hours, c.minutes-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
