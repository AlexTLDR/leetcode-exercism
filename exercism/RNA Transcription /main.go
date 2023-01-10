// https://exercism.org/tracks/go/exercises/rna-transcription

package main

import "fmt"

func main() {
	dna := "GCTA"
	fmt.Println(ToRNA(dna))
}

func ToRNA(dna string) string {
	conversion := ""
	DNAtoARN := map[rune]rune{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	for _, v := range dna {
		v = DNAtoARN[v]
		conversion += string(v)
	}
	return conversion
}
