// https://exercism.org/tracks/go/exercises/isogram

package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsIsogram("lumbserjacks"))
}

func IsIsogram(word string) bool {
	var wordMap map[rune]int = make(map[rune]int)
	for _, v := range word {
		wordMap[v]++
	}
	for _, v := range wordMap {
		if v > 1 {
			return false
		}
	}
	return true
}
