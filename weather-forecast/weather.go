// Package weather it says that the package name is weather.
package weather

var (
	// CurrentCondition should be equal to the condition.
	CurrentCondition string
	// CurrentLocation is bla bla where you are location.
	CurrentLocation string
)

// Forecast variable.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
