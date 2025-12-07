package utils

import "math"

func Round(val float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Round(val*shift) / shift
}
