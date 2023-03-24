// https://exercism.org/tracks/go/exercises/bob

package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func main() {
	fmt.Println(Hey("You are, what, like 15?"))
}

func Hey(remark string) string {
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	hasDecimal := false
	hasLower := false
	isQuestion := false
	if remark[len(remark)-1] == 63 {
		isQuestion = true
	}
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z]+`)
	remark = nonAlphanumericRegex.ReplaceAllString(remark, "")
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	for _, r := range remark {
		if unicode.IsLower(r) {
			hasLower = true
		}
		if unicode.IsDigit(r) {
			hasDecimal = true
		}
	}

	if isQuestion {
		if hasLower || hasDecimal {
			return "Sure."
		} else {
			return "Calm down, I know what I'm doing!"
		}
	}
	if hasDecimal {
		return "Whatever."
	}
	if !hasLower {
		return "Whoa, chill out!"
	}
	for _, r := range remark {
		if r == 32 {
			return "Fine. Be that way!"
		}
	}
	return "Whatever."
}
