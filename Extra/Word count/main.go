package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `
	dogs and cats
	dogs and cats
	cats and mice
	cats and mice
	mice eat cheese
	`

	words := strings.Fields(text)
	counts := map[string]int{}
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	fmt.Println(counts)
}
