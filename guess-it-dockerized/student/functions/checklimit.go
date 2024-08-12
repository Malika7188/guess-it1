package mathskils

import "math"

func ChecLimit(number []float64) (int, int) {

	start := len(number) - 3
	if start < 0 {
		start = 0
	}

	win := number[start:]
	if len(number) == 0 {
		return 0, 0
	}
	mean := HandleAverage(win)
	//fmt.Print(mean)
	std := HandleStandardDeviation(win)

	lowerLimit := mean - (3 * std)
	upperLimit := mean + (3 * std)

	return int(math.Round(lowerLimit)), int(math.Round(upperLimit))
}
