package strand

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
