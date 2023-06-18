// https://exercism.org/tracks/go/exercises/largest-series-product

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(LargestSeriesProduct("73167176531330624919225119674426574742355349194934", 3))
}

func LargestSeriesProduct2(digits string, span int) (int64, error) {
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
