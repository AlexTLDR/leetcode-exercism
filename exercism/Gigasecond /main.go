//https://exercism.org/tracks/go/exercises/gigasecond

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(AddGigasecond(time.Date(2046, time.October, 2, 23, 46, 40, 0, time.Local)))
}

func AddGigasecond(t time.Time) time.Time {
	//gigaSecond := 1000000000 * time.Second
	gigaSecond := 1e9 * time.Second
	return t.Add(gigaSecond)
}
