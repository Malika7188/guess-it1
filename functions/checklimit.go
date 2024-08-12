package mathskils

func ChecLimit(number []float64) (int64, int64) {
	if len(number) == 0 {
		return 0, 0
	}
	mean := HandleAverage(number)
	//fmt.Print(mean)
	std := HandleStandardDeviation(number)

	lowerLimit := int64(mean - (2 * std))
	upperLimit := int64(mean + (2 * std))

	return lowerLimit, upperLimit
}
