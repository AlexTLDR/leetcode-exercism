package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	wordCount := make(map[string]int)
	words := regexp.MustCompile("[a-zA-Z0-9]+(['][a-zA-Z0-9]+)?").FindAllString(phrase, -1)

	for _, w := range words {
		wordCount[strings.ToLower(w)]++
	}
	return wordCount
}
