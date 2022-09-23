package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	multipliers := []int{3, 2, 1}
	t := maximumScore(nums, multipliers)
	fmt.Println(t)
}

func maximumScore(nums []int, multipliers []int) int {
	//n, m := len(nums), len(multipliers)
	if nums[0]*multipliers[0] > nums[len(nums)-1]*multipliers[0] {
		return nums[0] * multipliers[0]
	} else {
		return nums[len(nums)-1] * multipliers[0]
	}

	// left := maximumScore(nums[1], multipliers[0])
	// right := maximumScore(nums[len(nums)-1], multiplies[len(multipliers)-1])
	// return max(left, right)
	//work in progress, not yet finished

}
