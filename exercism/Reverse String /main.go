// https://exercism.org/tracks/go/exercises/reverse-string

package main

import "fmt"

func main() {

	input := "cool"
	output := Reverse(input)
	fmt.Println(output)

}

func Reverse(input string) string {
	inputArr := []rune(input)
	for i, j := 0, len(inputArr)-1; i < j; i, j = i+1, j-1 {
		inputArr[i], inputArr[j] = inputArr[j], inputArr[i]
	}
	return string(inputArr)
}
