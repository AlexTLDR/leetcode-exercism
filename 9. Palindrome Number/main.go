package main

import "fmt"

func main() {
	x := 123454321
	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
	number := x
	var reversedNumber int
	for number > 0 {
		reversedNumber = reversedNumber*10 + number%10
		number = number / 10
	}
	return x == reversedNumber
}
