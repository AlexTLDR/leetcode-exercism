// https://exercism.org/tracks/go/exercises/booking-up-for-beauty

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Schedule("11/22/2022 9:00:00"))
}

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	appointment, err := time.Parse("1/02/2006 15:04:05", date) //"1/02/2006 15:04:05" the exact date used by go to parse time
	if err != nil {
		panic(err)
	}
	return appointment
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	panic("Please implement the HasPassed function")
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
