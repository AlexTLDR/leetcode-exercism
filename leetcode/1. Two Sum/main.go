package main

import "fmt"

func main() {
	result := twoSum([]int{3, 3}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	for i, v := range nums {
		for j, w := range nums {

			if v+w == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
