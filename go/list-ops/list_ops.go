package listops

// IntList is an array of ints
type IntList []int
type unaryFunc func(int) int
type binFunc func(int, int) int
type predFunc func(int) bool

// Foldr implementation
func (items IntList) Foldr(f binFunc, initial int) int {
	for i := len(items) - 1; i >= 0; i-- {
		initial = f(items[i], initial)
	}

	return initial
}

// Foldl implementation
func (items IntList) Foldl(f binFunc, initial int) int {
	for _, item := range items {
		initial = f(initial, item)
	}

	return initial
}

// Filter implementation
func (items IntList) Filter(f predFunc) IntList {
	var filtered IntList = make([]int, len(items))
	var counter = 0

	for _, item := range items {
		if f(item) {
			filtered[counter] = item
			counter++
		}
	}

	return filtered[:counter]
}

// Map implementation
func (items IntList) Map(f unaryFunc) IntList {
	var transformed IntList = make([]int, len(items))

	for i, item := range items {
		transformed[i] = f(item)
	}

	return transformed
}

// Length returns the length of the IntList
func (items IntList) Length() int {
	return len(items)
}

// Reverse implementation
func (items IntList) Reverse() IntList {
	var reversed IntList = make([]int, len(items))
	var total = len(items) - 1

	for i, item := range items {
		reversed[total-i] = item
	}

	return reversed
}

// Append implementation
func (items IntList) Append(others IntList) IntList {
	var index = len(items)
	var result IntList = make([]int, index+len(others))

	for i, item := range items {
		result[i] = item
	}

	for i, item := range others {
		result[index+i] = item
	}

	return result
}

// Concat implementation
func (items IntList) Concat(lists []IntList) IntList {
	for _, others := range lists {
		items = items.Append(others)
	}

	return items
}
