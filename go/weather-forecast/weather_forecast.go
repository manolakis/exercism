// Package weather provides information about the weather at a specific location.
package weather

// CurrentCondition contains the current weather condition.
var CurrentCondition string

// CurrentLocation contains the current weather location.
var CurrentLocation string

// Forecast returns a string describing the weather at a specific location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
