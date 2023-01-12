// https://exercism.org/tracks/go/exercises/anagram

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	target := "stone"
	candidates := []string{"stone", "tones", "banana", "tons", "notes", "Seton"}

	//the anagram set is "tones", "notes", "Seton"

	fmt.Println(Detect(target, candidates))

}

func Detect(subject string, candidates []string) []string {

	anagrams := []string{}
	subjectSlice := strings.Split(strings.ToLower(subject), "")
	sort.Slice(subjectSlice, func(i, j int) bool {
		return subjectSlice[i] < subjectSlice[j]
	})
	fmt.Println(subjectSlice)
	for _, v := range candidates {
		vSlice := strings.Split(strings.ToLower(v), "")
		sort.Slice(vSlice, func(i, j int) bool {
			return vSlice[i] < vSlice[j]
		})
		if SliceCompare(subjectSlice, vSlice) {
			anagrams = append(anagrams, v)
		}
	}

	return anagrams
}

func SliceCompare(subject, candidate []string) bool {
	if len(subject) != len(candidate) {
		return false
	}

	for i, v := range subject {
		if v != candidate[i] {
			return false
		}
	}
	return true
}
