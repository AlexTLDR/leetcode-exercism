// https://exercism.org/tracks/go/exercises/grains

package main

import "fmt"

const minsInDay = 24 * 60

type Clock struct {
	hour    int
	minutes int
}

func main() {
	fmt.Println(New(144, 60))
}

func New(h, m int) Clock {
	if m%60 == 0 {
		h = h + m/60
	}
	fmt.Println(h)
	h = h % 24
	fmt.Println(h)
	m = m % 60
	return Clock{h, m}
}

// func (c Clock) Add(m int) Clock {
// 	panic("Please implement the Add function")
// }

// func (c Clock) Subtract(m int) Clock {
// 	panic("Please implement the Subtract function")
// }

// func (c Clock) String() string {
// 	panic("Please implement the String function")
// }
