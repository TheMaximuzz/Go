package lasagna

import "fmt"

// TODO: define the 'OvenTime' constant
const OvenTime int = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    var tmpTime int = OvenTime
    tmpTime -= actualMinutesInOven
    fmt.Printf("first func answer is: %v\n", tmpTime)
    return tmpTime
}


// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	var fullTime int = numberOfLayers * 2
    fmt.Printf("sec func answer is: %v\n", fullTime)
    return fullTime
}


// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	var TotalTime int = numberOfLayers*2 + actualMinutesInOven
    fmt.Printf("Total time in the oven is %+v\n", TotalTime)
    return TotalTime
}