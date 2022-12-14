package main

import "fmt"

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools

type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

type Chessboard map[string]File

func main() {
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
	//A = []bool{true}
	//fmt.Printf("%v, %T\n", A, A)
	//cb := newChessboard()
	fmt.Println(CountInFile(board, "A"))
	//fmt.Println(CountInRank(board))

}

func newChessboard() Chessboard {
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
	for _, v := range cb[file] {
		if v {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	panic("Please implement CountInRank()")
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	panic("Please implement CountOccupied()")
}
