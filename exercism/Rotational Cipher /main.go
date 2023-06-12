// https://exercism.org/tracks/go/exercises/rotational-cipher

package main

import "fmt"

func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz"))
}

func RotationalCipher(plain string, shiftKey int) string {
	alphabet := map[int]rune{
		1:'a',
		2:'b',
		3:'c',
		4:'d',
		5:'e',
		6:'f',
		7:'g',
		8:'h',
		9:'i',
		10:'j',
		11:'k',
		12:'l',
		13:'m',
		14:'n',
		15:'o',
		16:'p',
		17:'q',
		18:'r',
		19:'s',
		20:'t',
		21:'u',
		22:'v',
		23:'w',
		24:'x',
		25:'y',
		26:'z',
	}

	for i, r := range plain {
		for s,j := range alphabet {
			if r == s {
plain[i] = 
			}
		}
	}
}
