// https://exercism.org/tracks/go/exercises/weather-forecast

// Package weather returns the current weather in the current location.
package weather

// CurrentCondition describes the current weather.
var CurrentCondition string

// CurrentLocation describes the current location (city).
var CurrentLocation string

// Forecast returns the current weather in the selected city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
