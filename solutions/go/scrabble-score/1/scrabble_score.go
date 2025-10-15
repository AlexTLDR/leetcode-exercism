package scrabble

import "strings"

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
