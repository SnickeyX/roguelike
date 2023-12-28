package utils

import "math"

// fov
func EuclidianDist(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(y2-y1), 2) + math.Pow(float64(x2-x1), 2))
}

// weapon range
func ChebyshevDist(x1, y1, x2, y2 int) float64 {
	return max(math.Abs(float64(x1-x2)), math.Abs(float64(y1-y2)))
}
