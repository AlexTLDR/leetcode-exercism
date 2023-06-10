// https://exercism.org/tracks/go/exercises/word-count

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(WordCount(`"That's the password: 'PASSWORD 123'!", cried the Special Agent.\nSo I fled.`))
}

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordCount := make(map[string]int)
	words := regexp.MustCompile("[a-zA-Z0-9]+(['][a-zA-Z0-9]+)?").FindAllString(phrase, -1)

	for _, w := range words {
		wordCount[strings.ToLower(w)]++
	}
	return wordCount
}
