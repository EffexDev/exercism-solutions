package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	f, _ := cb[file]
	count := 0
	for _, occupied := range f {
		if occupied {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	occupiedCount := 0
	index := rank - 1

	for _, file := range cb {
		if len(file) > index && file[index] {
			occupiedCount++
		}
	}
	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	totalSquares := 0
	for _, file := range cb {
		totalSquares += len(file)
	}
	return totalSquares
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupiedSquares := 0
	for _, file := range cb {
		for _, isOccupied := range file {
			if isOccupied {
				occupiedSquares++
			}
		}
	}
	return occupiedSquares
}
