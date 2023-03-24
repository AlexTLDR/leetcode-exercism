// https://exercism.org/tracks/go/exercises/bob

package main

import "fmt"

func main() {
	fmt.Println(Hey("fffbbcbeab?"))
}

func Hey(remark string) string {

	for _, r := range remark {
		if Capitals(r) != "lc" && remark[len(remark)-1] == 63 {
			return "Sure."
		}

		if r == 0 {
			return "Fine. Be that way!"
		}
	}

	for _, r := range remark {
		if Capitals(r) == "lc" {
			return "Whatever."
		}
	}
	for _, r := range remark {
		if Capitals(r) == "c" && remark[len(remark)-1] == 63 {
			return "Calm down, I know what I'm doing!"
		} else if Capitals(r) == "c" {
			return "Whoa, chill out!"
		}

	}
	return "Whatever."
}

func Capitals(i rune) string {
	switch {
	case i >= 97 && i <= 122:
		return "lc"
	case i >= 65 && i <= 90:
		return "c"
	}
	return "na"
}
