package lsproduct

import (
	"fmt"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	for _, v := range digits {
		if v < '0' || v > '9' {
			return 0, fmt.Errorf("input should only contain digits")
		}

	}
	if span <= 0 || span > len(digits) {
		return 0, fmt.Errorf("invalid span value")
	}

	var max int
	for i := 0; i < len(digits)-span+1; i++ {
		product := 1
		for j := i; j < i+span; j++ {
			product *= int(digits[j] - '0')
		}
		if product > max {
			max = product
		}
	}
	return int64(max), nil
}
