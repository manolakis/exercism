package darts

import "math"

func Score(x, y float64) int {
	distanceFromCenter := math.Sqrt(x*x + y*y)

	switch {
	case distanceFromCenter <= 1:
		return 10
	case distanceFromCenter <= 5:
		return 5
	case distanceFromCenter <= 10:
		return 1
	default:
		return 0
	}
}
