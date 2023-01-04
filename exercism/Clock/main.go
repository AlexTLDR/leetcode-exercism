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
	//c := Clock{23, 15}

	// fmt.Println(c.String())
	//fmt.Println(New(0, 160))

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
	// c.minutes -= m
	// if c.minutes < 0 {
	// 	fmt.Println("from subtract", c.hours, c.minutes)
	// 	c.hours -= 1 + c.minutes/60
	// 	fmt.Println(c.hours)
	// 	c.minutes = 60 + c.minutes%60
	// 	if c.minutes == 60 {
	// 		c.minutes = 0
	// 		c.hours += 1
	// 	}
	// }
	// if c.hours >= 24 {
	// 	c.hours %= 24
	// }
	return c.Add(-m)

}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hours, c.minutes)
}
