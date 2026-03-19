// Package weather provides tools for weather predict.
package weather

var (
    // CurrentCondition stores info about current weather conditions.
	CurrentCondition string
    // CurrentLocation stores info about current location.
	CurrentLocation  string
)

// Forecast returns info about weather. 
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
