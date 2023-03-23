//https://exercism.org/tracks/go/exercises/gigasecond

package main

import "time"

func main() {

}

func AddGigasecond(t time.Time) time.Time {
	gigaSecond := 1000000000 * time.Second
	return t.Add(gigaSecond)
}
