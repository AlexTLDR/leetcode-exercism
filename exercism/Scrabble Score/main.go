//https://exercism.org/tracks/go/exercises/scrabble-score

package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Score", Score("cabbage"))
	fmt.Println("ScoreMap:", ScoreMap("cabbage"))
}

func Score(word string) int {
	var onePoint = "AEIOULNRST"
	var twoPoints = "DG"
	var threePoints = "BCMP"
	var fourPoints = "FHVWY"
	var fivePoints = "K"
	var eightPoints = "JX"
	var tenPoints = "QZ"

	var totalScore int
	word = strings.ToUpper(word)
	for _, v := range word {
		if strings.ContainsRune(onePoint, v) {
			score := 1
			totalScore += score
		} else if strings.ContainsRune(twoPoints, v) {
			score := 2
			totalScore += score
		} else if strings.ContainsRune(threePoints, v) {
			score := 3
			totalScore += score
		} else if strings.ContainsRune(fourPoints, v) {
			score := 4
			totalScore += score
		} else if strings.ContainsRune(fivePoints, v) {
			score := 5
			totalScore += score
		} else if strings.ContainsRune(eightPoints, v) {
			score := 8
			totalScore += score
		} else if strings.ContainsRune(tenPoints, v) {
			score := 10
			totalScore += score
		}

	}

	return totalScore
}

func ScoreMap(word string) int {
	score := 0
	var values = map[rune]int{
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
		'L': 1,
		'N': 1,
		'R': 1,
		'S': 1,
		'T': 1,
		'D': 2,
		'G': 2,
		'B': 3,
		'C': 3,
		'M': 3,
		'P': 3,
		'F': 4,
		'H': 4,
		'V': 4,
		'W': 4,
		'Y': 4,
		'K': 5,
		'J': 8,
		'X': 8,
		'Q': 10,
		'Z': 10,
	}

	for _, l := range strings.ToUpper(word) {
		score += values[l]
	}
	return score
}
