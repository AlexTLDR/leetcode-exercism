// https://exercism.org/tracks/go/exercises/atbash-cipher

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Atbash("x123 yes"))
}

func Atbash(s string) string {
	var alphabetMap = map[string]string{
		"a": "z",
		"b": "y",
		"c": "x",
		"d": "w",
		"e": "v",
		"f": "u",
		"g": "t",
		"h": "s",
		"i": "r",
		"j": "q",
		"k": "p",
		"l": "o",
		"m": "n",
		"n": "m",
		"o": "l",
		"p": "k",
		"q": "j",
		"r": "i",
		"s": "h",
		"t": "g",
		"u": "f",
		"v": "e",
		"w": "d",
		"x": "c",
		"y": "b",
		"z": "a",
	}

	for _, v := range s {
		for w := range alphabetMap {
			if string(v) == w {
				s = strings.Replace(s, w, alphabetMap[w], 1)
			}
		}
	}
	return s

}
