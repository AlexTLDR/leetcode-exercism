// https://exercism.org/tracks/go/exercises/leap

package main

import "fmt"

func main() {
	fmt.Println(IsLeapYear(900))
}

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	}
	return false
}
