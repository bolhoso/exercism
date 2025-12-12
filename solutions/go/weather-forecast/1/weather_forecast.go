// Package weather has weather forecasting functions.
package weather

var (
    // CurrentCondition is the current weather condition.
	CurrentCondition string
    // CurrentLocation is the city name in the land of Goblins.
	CurrentLocation  string
)

// Forecast returns a string describing the current weather condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
