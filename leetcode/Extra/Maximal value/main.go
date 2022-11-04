package main

import "fmt"

func main() {
	nums := []int{23, 15, 3, 7, 27, 56}

	max := nums[0]

	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}
	fmt.Println(max)
}
