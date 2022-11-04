//https://exercism.org/tracks/go/exercises/scrabble-score

package main

import (
	"fmt"
	"strings"
)

func main() {
	// scores := make(map[int]string)
	// scores = map[int]string{
	// 	1 : "A",
	// 	2 :
	// 	3
	// 	4
	// 	5
	// 	8
	// 	10
	// }

	fmt.Println(Score("cabbage"))
	fmt.Println("taste:", Score("Taste"))
}

func Score(word string) int {
	var onePoint = "AEIOULNRST"
	var totalScore int
	word = strings.ToUpper(word)
	for _, v := range word {
		//if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' || v == 'l' || v == 'n' || v == 'r' || v == 's' || v == 't' {
		if strings.ContainsRune(onePoint, v) {
			score := 1
			totalScore += score
		}

	}

	return totalScore
}

// Letter                           Value
// A, E, I, O, U, L, N, R, S, T       1
// D, G                               2
// B, C, M, P                         3
// F, H, V, W, Y                      4
// K                                  5
// J, X                               8
// Q, Z                               10
