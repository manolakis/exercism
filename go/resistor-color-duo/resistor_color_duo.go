package resistorcolorduo

import "slices"

var COLORS = []string{"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"}

func colorCode(color string) int {
	return slices.Index(COLORS, color)
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	return colorCode(colors[0])*10 + colorCode(colors[1])
}
