package triangle

import (
	"math"
)

// Kind represents the triangle type
type Kind int

const (
	// NaT means not a triangle
	NaT Kind = 1 << iota
	// Equ means equilateral triangle
	Equ
	// Iso means isosceles triangle
	Iso
	// Sca means scalene triangle
	Sca
	// Deg means degenerate triangle
	Deg
)

// KindFromSides returns the kind of triangle.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !allSidesAreValid(a, b, c):
		return NaT
	case isDegenerate(a, b, c):
		return Deg
	case isInequal(a, b, c):
		return NaT
	case isEquilateral(a, b, c):
		return Equ
	case isIsosceles(a, b, c):
		return Iso
	case isScalene(a, b, c):
		return Sca
	default:
		return NaT
	}
}

func isEquilateral(a, b, c float64) bool {
	return a == b && b == c
}

func isIsosceles(a, b, c float64) bool {
	return !isEquilateral(a, b, c) && (a == b || b == c || a == c)
}

func isScalene(a, b, c float64) bool {
	return a != b && b != c && a != c
}

func allSidesAreValid(a, b, c float64) bool {
	return isValid(a) && isValid(b) && isValid(c)
}

func isValid(a float64) bool {
	return !math.IsNaN(a) && !math.IsInf(a, 0) && a > 0
}

func isDegenerate(a, b, c float64) bool {
	return a+b == c ||
		a+c == b ||
		b+c == a
}

func isInequal(a, b, c float64) bool {
	return a+b < c ||
		b+c < a ||
		a+c < b
}
