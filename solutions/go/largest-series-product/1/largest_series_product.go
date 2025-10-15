package lsproduct

import (
	"fmt"
	"strconv"
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

	var largestProduct int64 = 0

	for i := 0; i <= len(digits)-span; i++ {
		series := digits[i : i+span]
		var currentProduct int64 = 1

		for _, digit := range series {
			digitValue, _ := strconv.ParseInt(string(digit), 10, 64)
			currentProduct *= digitValue
		}

		if currentProduct > largestProduct {
			largestProduct = currentProduct
		}
	}

	return largestProduct, nil
}
