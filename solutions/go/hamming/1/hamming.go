package hamming

import "errors"

func Distance(a, b string) (int, error) {
	hammingDistance := 0
	if len(a) != len(b) {
		return -1, errors.New("DNA strands not of equal")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			hammingDistance += 1
		}
	}
	return hammingDistance, nil

}
