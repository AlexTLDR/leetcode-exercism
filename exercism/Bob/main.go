// https://exercism.org/tracks/go/exercises/bob

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Hey(""))
}

func Hey(remark string) string {
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	remark = strings.Replace(remark, " ", "", -1)
	remark = strings.Replace(remark, "\t", "", -1)
	remark = strings.Replace(remark, "\n", "", -1)
	remark = strings.Replace(remark, "\r", "", -1)
	// hasUpper := false
	hasLower := false
	for _, r := range remark {
		if unicode.IsLower(r) {
			hasLower = true
		}
	}
	if remark[len(remark)-1] == 63 {
		if hasLower {
			return "Sure."
		} else {
			return "Calm down, I know what I'm doing!"
		}
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
