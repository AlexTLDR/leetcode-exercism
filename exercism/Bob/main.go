// https://exercism.org/tracks/go/exercises/bob

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(Hey("TBBBBBBAAA?"))
}

func Hey(remark string) string {
	// if len(remark) == 0 {
	// 	return "Fine. Be that way!"
	// }
	// for _, r := range remark {
	// 	if r == 32 {
	// 		return "Fine. Be that way!"
	// 	}
	// 	if unicode.IsLower(r) && remark[len(remark)-1] == 63 {
	// 		return "Sure."
	// 	}
	// }

	// for _, r := range remark {
	// 	if unicode.IsLower(r) {
	// 		return "Whatever."
	// 	}
	// }
	// for _, r := range remark {
	// 	if unicode.IsUpper(r) && remark[len(remark)-1] == 63 {
	// 		return "Calm down, I know what I'm doing!"
	// 	} else if unicode.IsUpper(r) {
	// 		return "Whoa, chill out!"
	// 	}

	// }
	// return "Whatever."
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

	return ""
}
