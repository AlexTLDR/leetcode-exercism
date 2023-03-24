// https://exercism.org/tracks/go/exercises/bob

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Hey("1, 2, 3, 4"))
}

func Hey(remark string) string {
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	remark = strings.Replace(remark, " ", "", -1)
	remark = strings.Replace(remark, "\t", "", -1)
	remark = strings.Replace(remark, "\n", "", -1)
	remark = strings.Replace(remark, "\r", "", -1)
	hasDecimal := false
	hasLower := false
	for _, r := range remark {
		if unicode.IsLower(r) {
			hasLower = true
		}
		if unicode.IsDigit(r) {
			hasDecimal = true
		}
	}

	if remark[len(remark)-1] == 63 {
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
