package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordCount := make(map[string]int)
	words := regexp.MustCompile(`\w+('[A-z]+)?`).FindAllString(phrase, -1)

	for _, w := range words {
		wordCount[strings.ToLower(w)]++
	}
	return wordCount
}
