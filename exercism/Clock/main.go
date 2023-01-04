// https://exercism.org/tracks/go/exercises/grains

package main

import "fmt"

type Clock struct {
	hours   int
	minutes int
}

func main() {

	c := New(0, 10)
	fmt.Println(c)
	//fmt.Println(c.Add(50))
	fmt.Println(c.Add(-20))
	fmt.Println(c.Subtract(20))
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
