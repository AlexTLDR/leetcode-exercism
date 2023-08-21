// https://exercism.org/tracks/go/exercises/simple-cipher

package cipher

import (
	"strings"
	"unicode"
)

type ShiftCipher struct {
	length int
	shift  []int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return ShiftCipher{1, []int{distance}}
}

func NewVigenere(key string) Cipher {
	length := len(key)
	if length == 0 {
		return nil
	}
	shift := []int{}
	notZero := false
	for _, r := range key {
		if r < 'a' || r > 'z' {
			return nil
		}
		if r > 'a' {
			notZero = true
		}
		shift = append(shift, int(r-'a'))
	}
	if notZero {
		return ShiftCipher{length, shift}
	}
	return nil
}

func (sc ShiftCipher) Encode(plaintext string) string {
	var ciphertext strings.Builder
	idx := 0
	for _, r := range plaintext {
		if r < 'A' || (r > 'Z' && r < 'a') || r > 'z' {
			continue
		}
		encoder := (26 + sc.shift[idx])
		r = 'a' + (unicode.ToLower(r)-'a'+rune(encoder))%26
		ciphertext.WriteRune(r)
		idx = (idx + 1) % sc.length
	}
	return ciphertext.String()
}

func (sc ShiftCipher) Decode(ciphertext string) string {
	var plaintext strings.Builder
	idx := 0
	for _, r := range ciphertext {
		if r < 'a' || r > 'z' {
			continue
		}
		decoder := (26 - sc.shift[idx])
		r = 'a' + (unicode.ToLower(r)-'a'+rune(decoder))%26
		plaintext.WriteRune(r)
		idx = (idx + 1) % sc.length
	}
	return plaintext.String()
}
