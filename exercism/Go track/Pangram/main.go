// https://exercism.org/tracks/go/exercises/pangram

package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence1 := "The quick brown fox jumps over the lazy dog."
	sentence2 := "Not a pangram Not a pangram Not a pangram Not a pangram"
	fmt.Println(IsPangram(sentence1))
	fmt.Println(IsPangram(sentence2))
}

func IsPangram(input string) bool {

	alphabet := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
		'h': 0,
		'i': 0,
		'j': 0,
		'k': 0,
		'l': 0,
		'm': 0,
		'n': 0,
		'o': 0,
		'p': 0,
		'q': 0,
		'r': 0,
		's': 0,
		't': 0,
		'u': 0,
		'v': 0,
		'w': 0,
		'x': 0,
		'y': 0,
		'z': 0,
	}

	if len(input) < len(alphabet) {
		return false
	}
	input = strings.ToLower(input)
	for _, v := range input {
		if _, ok := alphabet[rune(v)]; ok {
			alphabet[rune(v)]++
		}
	}

	for _, l := range alphabet {
		if l == 0 {
			return false
		}
	}

	return true
}
