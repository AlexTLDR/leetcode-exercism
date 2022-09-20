package main

import "fmt"

func main() {
	nums := []int{3, 1, 4}
	multipliers := []int{2}
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

}
