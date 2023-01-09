// https://exercism.org/tracks/go/exercises/triangle

package main

import "fmt"

func main() {
	fmt.Println(KindFromSides(0, 10, 10))
}

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if a <= 0 || b <= 0 || c <= 0 {
		k = NaT
	}

	if a == b && b == c && a == c {
		k = Equ
	}

	if a == b || a == c || b == c {
		k = Iso
	}

	if a != b && a != c && b != c {
		k = Sca
	}

	return k
}
