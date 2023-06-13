// https://exercism.org/tracks/go/exercises/rotational-cipher

package main

import "fmt"

func main() {
	fmt.Println(RotationalCipher("abcdefghijklmnopqrstuvwxyz", 13))
}

func RotationalCipher(plain string, shiftKey int) string {
	// alphabet := map[int]rune{
	// 	1:'a',
	// 	2:'b',
	// 	3:'c',
	// 	4:'d',
	// 	5:'e',
	// 	6:'f',
	// 	7:'g',
	// 	8:'h',
	// 	9:'i',
	// 	10:'j',
	// 	11:'k',
	// 	12:'l',
	// 	13:'m',
	// 	14:'n',
	// 	15:'o',
	// 	16:'p',
	// 	17:'q',
	// 	18:'r',
	// 	19:'s',
	// 	20:'t',
	// 	21:'u',
	// 	22:'v',
	// 	23:'w',
	// 	24:'x',
	// 	25:'y',
	// 	26:'z',
	// }

	result := ""
	for _, char := range plain {
		if char >= 'a' && char <= 'z' {
			result += string(int(char) + shiftKey)
		} else if char >= 'A' && char <= 'Z' {
			result += string(int(char) + shiftKey)
		} else {
			result += string(char)
		}
	}
	return result

}
