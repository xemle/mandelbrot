package fractal

import (
	"fmt"
	"math/big"
)

// Complex implements complex numbers
type Complex struct {
	Real, Imaginary *big.Float
}

// NewComplex creates a new complex number
func NewComplex() *Complex {
	r := new(big.Float)
	i := new(big.Float)

	return &Complex{r, i}
}
func (z *Complex) String() string {
	return fmt.Sprintf("%s+%si", z.Real.String(), z.Imaginary.String())
}

// SetFloat64 sets float64 value
func (z *Complex) SetFloat64(r, i float64) *Complex {
	z.Real.SetFloat64(r)
	z.Imaginary.SetFloat64(i)
	return z
}

// SetPrec Set precition of complex number
func (z *Complex) SetPrec(prec uint) *Complex {
	z.Real.SetPrec(prec)
	z.Imaginary.SetPrec(prec)
	return z
}

// Set sets a complex number
func (z *Complex) Set(x *Complex) *Complex {
	z.Real.Set(x.Real)
	z.Imaginary.Set(x.Imaginary)
	return z
}

// Add adds tow complex numbers
func (z *Complex) Add(x, y *Complex) *Complex {
	r := new(big.Float).Add(x.Real, y.Real)
	i := new(big.Float).Add(x.Imaginary, y.Imaginary)

	z.Real.Set(r)
	z.Imaginary.Set(i)

	return z
}

// Mul multiplies tow complex numbers
func (z *Complex) Mul(x, y *Complex) *Complex {
	r1 := new(big.Float).Mul(x.Real, y.Real)
	r2 := new(big.Float).Mul(x.Imaginary, y.Imaginary)
	i1 := new(big.Float).Mul(x.Real, y.Imaginary)
	i2 := new(big.Float).Mul(x.Imaginary, y.Real)

	z.Real.Set(new(big.Float).Sub(r1, r2))
	z.Imaginary.Set(new(big.Float).Add(i1, i2))

	return z
}

// Square returns the square of given complex number
func (z *Complex) Square(x *Complex) *Complex {
	// = x*x
	// = (xr + xi) * (xr + xi)
	// = (xr*xr - xi*xi) + (xr*xi + xi*xr)i
	zr2 := new(big.Float).Mul(x.Real, x.Real)
	zi2 := new(big.Float).Mul(x.Imaginary, x.Imaginary)

	zri := new(big.Float).Mul(x.Real, x.Imaginary)

	z.Real.Set(zr2.Sub(zr2, zi2))
	z.Imaginary.Set(zri.Add(zri, zri))

	return z
}
