package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	multipliers := []int{3, 2, 1}
	t := maximumScore(nums, multipliers)
	fmt.Println(t)
}

func maximumScore(nums []int, multipliers []int) int {
	if len(multipliers) == 1 {
		if nums[0]*multipliers[0] > nums[len(nums)-1]*multipliers[0] {
			return nums[0] * multipliers[0]
		} else {
			return nums[len(nums)-1] * multipliers[0]
		}
	}

	left := nums[0]*multipliers[0] + maximumScore(nums[1:], multipliers[1:])
	right := nums[len(nums)-1]*multipliers[0] + maximumScore(nums[:len(nums)-1], multipliers[1:])
	return max(left, right)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
