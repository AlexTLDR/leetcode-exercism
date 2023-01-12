// https://exercism.org/tracks/go/exercises/protein-translation

package main

import (
	"fmt"
)

var codon = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Phenylalanine",
	"UCC": "Phenylalanine",
	"UCA": "Phenylalanine",
	"UCG": "Phenylalanine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func main() {
	RNA := "AUGUUUUCU"
	fmt.Println(FromRNA(RNA))
}

func FromRNA(rna string) ([]string, error) {
	rnaSlice := []string{}
	// res := ""
	// for i, r := range rna {
	// 	res = res + string(r)
	// 	fmt.Printf("i %d r %c\n", i, r)
	// 	if i > 0 && (i+1)%3 == 0 {
	// 		fmt.Printf("=>(%d) '%v'\n", i, res)
	// 	}
	// }

	for i, r := range rna {
		tmp := ""
		if (i+1)%3 == 0 {
			tmp += string(r)
			fmt.Println(tmp)
		}
	}

	return rnaSlice, fmt.Errorf("RNA invalid")
}

func FromCodon(codon string) (string, error) {
	panic("Please implement the FromCodon function")
}
