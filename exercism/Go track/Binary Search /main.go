// https://exercism.org/tracks/go/exercises/binary-search

package main

func SearchInts(list []int, key int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2

		if list[mid] < key {
			low = mid + 1
		} else if list[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
