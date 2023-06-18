// https://exercism.org/tracks/go/exercises/phone-number

package main

import "fmt"

func main() {
	phoneNumber := "+1 (613)-995-0253"
	fmt.Println(Number(phoneNumber))
	fmt.Println(AreaCode(phoneNumber))
	fmt.Println(Format(phoneNumber))

}

func Number(phoneNumber string) (string, error) {
	panic("Please implement the Number function")
}

func AreaCode(phoneNumber string) (string, error) {
	panic("Please implement the AreaCode function")
}

func Format(phoneNumber string) (string, error) {
	panic("Please implement the Format function")
}
