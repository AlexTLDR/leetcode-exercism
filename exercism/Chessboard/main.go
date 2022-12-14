package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

func main() {
	/*
		these are for me to better understand the structure

		var A File = []bool{true, false, true, false, false, false, false, true}
		var B File = []bool{false, false, false, false, true, false, false, false}
		var C File = []bool{false, false, true, false, false, false, false, false}
		var D File = []bool{false, false, false, false, false, false, false, false}
		var E File = []bool{false, false, false, false, false, true, false, true}
		var F File = []bool{false, false, false, false, false, false, false, false}
		var G File = []bool{false, false, false, true, false, false, false, false}
		var H File = []bool{true, true, true, true, true, true, false, true}
		board := map[string]File{
			"A": A,
			"B": B,
			"C": C,
			"D": D,
			"E": E,
			"F": F,
			"G": G,
			"H": H,
		}

		using the NewChessboard func to generate the board as a more elegant way
	*/
	board := NewChessboard()
	fmt.Println(CountInFile(board, "A"))
	fmt.Println(CountInRank(board, 2))
	fmt.Println(CountAll(board))
	fmt.Println(CountOccupied(board))
}

func NewChessboard() Chessboard {
	return Chessboard{
		"A": File{true, false, true, false, false, false, false, true},
		"B": File{false, false, false, false, true, false, false, false},
		"C": File{false, false, true, false, false, false, false, false},
		"D": File{false, false, false, false, false, false, false, false},
		"E": File{false, false, false, false, false, true, false, true},
		"F": File{false, false, false, false, false, false, false, false},
		"G": File{false, false, false, true, false, false, false, false},
		"H": File{true, true, true, true, true, true, false, true},
	}
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	for _, element := range cb[file] {
		if element {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return count
	}

	for _, element := range cb {
		if element[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, element := range cb {

		count += len(element)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for key := range cb {
		count += CountInFile(cb, key)
	}
	return count
}
