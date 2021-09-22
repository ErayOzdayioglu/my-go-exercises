package lasagna

// TODO: define the 'OvenTime' constant

const OvenTime int = 40

// TODO: define the 'RemainingOvenTime()' function

func RemainingOvenTime(time int) int {
	return (OvenTime - time)
}

// TODO: define the 'PreparationTime()' function

func PreparationTime(count int) int {
	return count * 2
}

// TODO: define the 'ElapsedTime()' function

func ElapsedTime(count, time int) int {
	return time + PreparationTime(count)
}
