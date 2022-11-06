package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsIsogram("slumberjacks"))
}

func IsIsogram(word string) bool {
	for i := len(word) - 1; i > 0; i-- {
		for j := 1; j < i; j++ {
			fmt.Println("i:", word[i], "j:", word[j])
		}
		return false
	}
	return true
}
