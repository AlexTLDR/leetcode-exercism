// https://exercism.org/tracks/go/exercises/anagram

package main

import "fmt"

func main() {
	target := "stone"
	candidates := []string{"stone", "tones", "banana", "tons", "notes", "Seton"}

	//the anagram set is "tones", "notes", "Seton"

	fmt.Println(Detect(target, candidates))

}

func Detect(subject string, candidates []string) []string {
	panic("Please implement the Detect function")
}
