package protein

import (
	"errors"
)

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrInvalidBase = errors.New("ErrInvalidBase")
var ErrStop = errors.New("ErrStop")

func FromRNA(rna string) ([]string, error) {
	rnaSlice := []string{}
	for i := range rna {
		if (i+1)%3 == 0 {
			tmp, ok := FromCodon(string(rna[i-2]) + string(rna[i-1]) + string(rna[i]))
			if ok == ErrInvalidBase {
				return rnaSlice, ErrInvalidBase
			} else if ok == ErrStop {
				return rnaSlice, nil
			}
			rnaSlice = append(rnaSlice, tmp)
		}
	}

	return rnaSlice, nil
}

func FromCodon(codon string) (string, error) {
	for k, v := range codonMap {
		if k == codon {
			if v == "STOP" {
				return v, ErrStop
			}
			return v, nil
		}
	}
	return "", ErrInvalidBase
}