// https://exercism.org/tracks/go/exercises/series

package main

import "fmt"

func main() {
	fmt.Println(UnsafeFirst(3, "123456"))
}

func All(n int, s string) []string {
	panic("Please implement the All function")
}

func UnsafeFirst(n int, s string) string {
	if n > len(s) {
		return ""
	}
	var rs string
	for i := 0; i < n; i++ {
		rs += string(s[i])
	}
	return rs
}
