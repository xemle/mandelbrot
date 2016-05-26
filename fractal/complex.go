package fractal

import "fmt"

// Complex implements complex numbers
type Complex struct {
	Real, Imaginary float64
}

// NewComplex creates a new complex number
func NewComplex() *Complex {
	return &Complex{}
}

// NewComplex64 initialize a complex number with float64
func NewComplex64(r, i float64) *Complex {
	return &Complex{r, i}
}

// String converts the complex number to string
func (z *Complex) String() string {
	var sign = "+"
	if z.Imaginary < 0 {
		sign = ""
	}

	return fmt.Sprintf("%f%s%fi", z.Real, sign, z.Imaginary)
}

// Set sets a complex number
func (z *Complex) Set(x *Complex) *Complex {
	z.Real = x.Real
	z.Imaginary = x.Imaginary
	return z
}

// Add adds tow complex numbers
func (z *Complex) Add(x, y *Complex) *Complex {
	z.Real = x.Real + y.Real
	z.Imaginary = x.Imaginary + y.Imaginary

	return z
}

// Mul multiplies tow complex numbers
func (z *Complex) Mul(x, y *Complex) *Complex {
	r1 := x.Real * y.Real
	r2 := x.Imaginary * y.Imaginary
	i1 := x.Real * y.Imaginary
	i2 := x.Imaginary * y.Real

	z.Real = r1 - r2
	z.Imaginary = i1 + i2

	return z
}

// Square returns the square of given complex number
func (z *Complex) Square(x *Complex) *Complex {
	// = x*x
	// = (xr + xi) * (xr + xi)
	// = (xr*xr - xi*xi) + (xr*xi + xi*xr)i
	zr2 := x.Real * x.Real
	zi2 := x.Imaginary * x.Imaginary

	zri := x.Real * x.Imaginary

	z.Real = zr2 - zi2
	z.Imaginary = zri + zri

	return z
}
