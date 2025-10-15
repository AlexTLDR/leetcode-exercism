package meetup

import "time"

const testVersion = 3

type WeekSchedule int

const (
	First  = WeekSchedule(iota)
	Second
	Third
	Fourth
	Last
	Teenth
)

// Day returns the calendar day that matches the meetup restrictions
func Day(schedule WeekSchedule, day time.Weekday, month time.Month, year int) int {
	date := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	matches := []int{}
	dayDifference := (int(day) - int(date.Weekday()) + 7) % 7

	if dayDifference != 0 {
		date = date.AddDate(0, 0, dayDifference)
	}

	for date.Month() == month {
		matches = append(matches, date.Day())
		date = date.AddDate(0, 0, 7)
	}

	switch schedule {
	case First:
		return matches[0]
	case Second:
		return matches[1]
	case Third:
		return matches[2]
	case Fourth:
		return matches[3]
	case Last:
		return matches[len(matches)-1]
	case Teenth:
		for _, d := range matches {
			if d >= 13 {
				return d
			}
		}
	}

	panic("Invalid schedule")
}