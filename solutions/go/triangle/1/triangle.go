package triangle


type Kind string

const (

	NaT = "not a triangle"
	Equ = "equilateral"
	Iso = "isosceles"
	Sca = "scalene"
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	switch {
	case (a <= 0 || b <= 0 || c <= 0):
		k = NaT
	case a+b < c || a+c < b || b+c < a:
		k = NaT
	case a == b && b == c && a == c:
		k = Equ
	case a == b || a == c || b == c:
		k = Iso
	case a != b && a != c && b != c:
		k = Sca
	}

	return k
}
