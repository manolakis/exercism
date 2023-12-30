package chessboard

type File []bool

type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(chessboard Chessboard, file string) int {
	counter := 0

	for _, value := range chessboard[file] {
		if value {
			counter++
		}
	}

	return counter
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(chessboard Chessboard, rank int) int {
	counter := 0

	if rank >= 1 && rank <= 8 {
		for _, file := range chessboard {
			if file[rank-1] {
				counter++
			}
		}
	}

	return counter
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(chessboard Chessboard) int {
	counter := 0

	for _, file := range chessboard {
		counter += len(file)
	}

	return counter
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(chessboard Chessboard) int {
	counter := 0

	for file := range chessboard {
		counter += CountInFile(chessboard, file)
	}

	return counter
}
