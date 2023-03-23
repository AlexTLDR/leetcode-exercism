// https://exercism.org/tracks/go/exercises/bob

package main

import "fmt"

func main() {
	fmt.Println(Hey("hEy BOB?"))
}

func Hey(remark string) string {

	for _, r := range remark {
		if Capitals(r) == "c" && remark[len(remark)-1] == 63 {
			return "Sure."
		}
		if Capitals(r) == "lc" && remark[len(remark)-1] == 63 {
			return "Whoa, chill out!"
		}
	}
	return "Whatever."
}

func Capitals(i rune) string {
	switch {
	case i >= 97 && i <= 122:
		return "c"
	case i >= 65 && i <= 90:
		return "lc"
	}
	return "na"
}
