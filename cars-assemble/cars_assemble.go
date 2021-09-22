package cars

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return successRate(speed) * 221 * float64(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	var rate float64

	if speed == 0 {
		rate = 0.0
	} else if speed >= 1 && speed <= 4 {
		rate = 1.0
	} else if speed >= 5 && speed <= 8 {
		rate = 0.90
	} else {
		rate = 0.77
	}
	return rate
}
