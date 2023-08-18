// https://exercism.org/tracks/go/exercises/meetup

package main

import "time"

type WeekSchedule int

const (
	First  WeekSchedule = 1
	Second WeekSchedule = 2
	Third  WeekSchedule = 3
	Fourth WeekSchedule = 4
	Last   WeekSchedule = 5
	Teenth WeekSchedule = 6
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	var day int
	switch wSched {
	case First:
		day = 1
	case Second:
		day = 8
	case Third:
		day = 15
	case Fourth:
		day = 22
	case Last:
		day = lastDay(month, year)
	case Teenth:
		day = 13
	}
	return day + int(wDay-time.Date(year, month, 1, 0, 0, 0, 0, time.UTC).Weekday()+7)%7
}
