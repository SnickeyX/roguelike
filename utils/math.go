package utils

import "math"

func Min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func EuclidianDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(y2-y1), 2) + math.Pow(float64(x2-x1), 2))
}
