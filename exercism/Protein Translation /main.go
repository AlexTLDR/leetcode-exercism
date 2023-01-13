// https://exercism.org/tracks/go/exercises/protein-translation

package main

import (
	"errors"
	"fmt"
)

var codonMap = map[string]string{
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

	codon := "UUU"
	fmt.Println(FromCodon(codon))
}

func FromRNA(rna string) ([]string, error) {
	rnaSlice := []string{}
	err := errors.New("RNA invalid")
	for i := range rna {
		tmp := ""
		if (i+1)%3 == 0 {
			tmp = string(rna[i-2]) + string(rna[i-1]) + string(rna[i])
			rnaSlice = append(rnaSlice, RNAtoCodon(tmp))
		}
		err = nil
	}

	return rnaSlice, err
}

func FromCodon(codon string) (string, error) {
	for k, v := range codonMap {
		if k == codon {
			return v, nil
		}
	}
	return "", errors.New("is not a valid codon")
}

func RNAtoCodon(codon string) string {
	for k, v := range codonMap {
		if k == codon {
			return v
		}
	}
	return ""
}
