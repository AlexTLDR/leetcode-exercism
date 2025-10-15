package complexnumbers

import "math"

type Number struct {
	real      float64
	imaginary float64
}

func (n Number) Real() float64 {
	return n.real
}

func (n Number) Imaginary() float64 {
	return n.imaginary
}

func (n1 Number) Add(n2 Number) Number {
	return Number{
		real:      n1.real + n2.real,
		imaginary: n1.imaginary + n2.imaginary,
	}
}

func (n1 Number) Subtract(n2 Number) Number {
	return Number{
		real:      n1.real - n2.real,
		imaginary: n1.imaginary - n2.imaginary,
	}
}

func (n1 Number) Multiply(n2 Number) Number {
	r1, i1, r2, i2 := n1.real, n1.imaginary, n2.real, n2.imaginary
	return Number{
		real:      r1*r2 - i1*i2,
		imaginary: i1*r2 + r1*i2,
	}
}

func (n Number) Times(factor float64) Number {
	return Number{
		real:      factor * n.real,
		imaginary: factor * n.imaginary,
	}
}

func (n1 Number) Divide(n2 Number) Number {
	r1, i1, r2, i2 := n1.real, n1.imaginary, n2.real, n2.imaginary
	divisor := r2*r2 + i2*i2
	return Number{
		real:      (r1*r2 + i1*i2) / divisor,
		imaginary: (i1*r2 - r1*i2) / divisor,
	}
}

func (n Number) Conjugate() Number {
	return Number{
		real:      n.real,
		imaginary: -n.imaginary,
	}
}

func (n Number) Abs() float64 {
	return math.Sqrt(n.real*n.real + n.imaginary*n.imaginary)
}

func (n Number) Exp() Number {
	expReal := math.Exp(n.real)
	return Number{
		real:      expReal * math.Cos(n.imaginary),
		imaginary: expReal * math.Sin(n.imaginary),
	}
}