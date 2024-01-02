package diffsquares

func square(number int) int {
	return number * number
}

func reduceNaturalNumbers(numbers int, operation func(int) int) int {
	var acc int

	for i := 1; i <= numbers; i++ {
		acc += operation(i)
	}

	return acc
}

func SquareOfSum(numbers int) int {
	return square(reduceNaturalNumbers(numbers, func(number int) int {
		return number
	}))
}

func SumOfSquares(numbers int) int {
	return reduceNaturalNumbers(numbers, func(number int) int {
		return square(number)
	})
}

func Difference(numbers int) int {
	return SquareOfSum(numbers) - SumOfSquares(numbers)
}
