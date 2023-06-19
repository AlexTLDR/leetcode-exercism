// https://exercism.org/tracks/go/exercises/series

package main

import "fmt"

func main() {
	fmt.Println(UnsafeFirst(3, "123456"))
	fmt.Println(All(3, "1234567"))
}

func All(n int, s string) []string {
	var rs []string

	for i := 0; i < len(s)-n; i++ {
		ss := s[i : i+n]
		rs = append(rs, ss)
	}
	return rs
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
