// https://exercism.org/tracks/go/exercises/reverse-string

package main

import "fmt"

func main() {

	input := "cool"
	output := Reverse(input)
	fmt.Println(output)

}

func Reverse(input string) string {
	return input
}
