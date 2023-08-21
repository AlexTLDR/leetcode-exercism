// https://exercism.org/tracks/go/exercises/isogram

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(IsIsogram("lumbee r ja-c-ks"))
}

func IsIsogram(word string) bool {
	var wordMap map[rune]bool = make(map[rune]bool)
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")
	for _, v := range strings.ToLower(word) {

		if wordMap[v] == true {
			return false
		}
		wordMap[v] = true
	}
	fmt.Println("from println:", wordMap)
	// for _, v := range wordMap {
	// 	if v > 1 {
	// 		return false
	// 	}
	// }
	return true
}
