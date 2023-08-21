// https://exercism.org/tracks/go/exercises/grains

package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func main() {

	c := New(27, 80)
	fmt.Println("New: ", c)
	fmt.Println("Add: ", c.Add(50))
	fmt.Println("Subtract: ", c.Subtract(120))
	fmt.Println("Display: ", c.String())
}

func New(h, m int) Clock {
	h += m / 60
	m = m % 60
	if m < 0 {
		m += 60
		h -= 1
	}
	h = h % 24
	if h < 0 {
		h += 24
	}
	return Clock{h, m}
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
