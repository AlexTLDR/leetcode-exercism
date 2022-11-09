// https://exercism.org/tracks/go/exercises/isogram

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsIsogram("lumbserjacks"))
}

func IsIsogram(word string) bool {
	var wordMap map[rune]int = make(map[rune]int)
	lowerCaseWord := strings.ToLower(word)
	lowerCaseWord = strings.Replace(lowerCaseWord, "-", "", -1)
	lowerCaseWord = strings.Replace(lowerCaseWord, " ", "", -1)
	for _, v := range lowerCaseWord {
		wordMap[v]++
	}
	for _, v := range wordMap {
		if v > 1 {
			return false
		}
	}
	return true
}
