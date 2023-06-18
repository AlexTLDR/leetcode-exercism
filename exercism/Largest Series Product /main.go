// https://exercism.org/tracks/go/exercises/largest-series-product

package main

import "fmt"

func main() {
	fmt.Println(LargestSeriesProduct("63915", 3))
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span || span < 2 {
		return 0, fmt.Errorf("invalid")
	}
	var max int
	for i := 0; i < len(digits)-span+1; i++ {
		var product int
		for j := i; j < i+span; j++ {
			product *= int(digits[j] - '0')
		}
		if product > max {
			max = product
		}
	}
	return int64(max), nil
}
