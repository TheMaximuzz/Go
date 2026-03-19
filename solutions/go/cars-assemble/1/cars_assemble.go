package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var res float64 =  float64(productionRate) * (successRate*0.01)
    return res
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var res float64 =  float64(productionRate) * (successRate*0.01)
    var ans int = int(res)/60
    return ans
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var oneCars int = carsCount % 10
	var tenCars int = carsCount / 10
    var res uint =  uint(tenCars*95000 + oneCars*10000)
    return res
}
