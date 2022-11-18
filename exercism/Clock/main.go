// https://exercism.org/tracks/go/exercises/grains

package main

import "fmt"

const minsInDay = 24 * 60

type Clock struct {
	hours   int
	minutes int
}

func main() {
	// fmt.Println(New(144, 60))
	// c := Clock{23, 15}
	// fmt.Println(c.Add(50))
	// fmt.Println(c.String())
	fmt.Println(New(0, 160))
}

func New(h, m int) Clock {
	// if m%60 == 0 {
	// 	h = h + m/60
	// }
	// fmt.Println(h)
	// h %= 24
	// fmt.Println(h)
	// m %= 60
	// return Clock{h, m}
	clock := Clock{0, 0}
	clock = clock.Add(h*60 + m)
	return clock

}

func (c Clock) Add(m int) Clock {
	c.minutes += m
	if c.minutes >= 60 {
		c.hours += c.minutes / 60
		c.minutes %= 60
	}
	if c.hours >= 24 {
		c.hours %= 24
	}
	return c
}

func (c Clock) Subtract(m int) Clock {
	panic("Please implement the Subtract function")
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
