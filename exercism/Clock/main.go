// https://exercism.org/tracks/go/exercises/grains

package main

import "fmt"

const minsInDay = 24 * 60

type Clock int

func main() {
	fmt.Println(New(24, 35))
}

func New(h, m int) Clock {
	clockTime := h*60 + m
	if clockTime <= minsInDay {
		return Clock(clockTime)
	}
	return Clock(clockTime) - minsInDay
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
