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
	word = strings.Replace(word, "-", "", -1)
	word = strings.Replace(word, " ", "", -1)
	for _, v := range strings.ToLower(word) {
		wordMap[v]++
	}
	for _, v := range wordMap {
		if v > 1 {
			return false
		}
	}
	return true
}
