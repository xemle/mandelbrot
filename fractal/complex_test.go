package fractal_test

import (
	"mandelbrot/fractal"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexToString(t *testing.T) {
	s := fractal.NewComplex()
	assert.Equal(t, "0+0i", s.String())
}

func TestSetFloat64(t *testing.T) {
	z := fractal.NewComplex().SetFloat64(1, 0)
	r := new(big.Float).SetFloat64(1)
	i := new(big.Float).SetFloat64(0)
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestAdd(t *testing.T) {
	x := fractal.NewComplex().SetFloat64(1, 0)
	y := fractal.NewComplex().SetFloat64(0, 2)
	z := fractal.NewComplex().Add(x, y)

	r := new(big.Float).SetFloat64(1)
	i := new(big.Float).SetFloat64(2)
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestMul(t *testing.T) {
	x := fractal.NewComplex().SetFloat64(1, 2)
	y := fractal.NewComplex().SetFloat64(3, -1)
	z := fractal.NewComplex().Mul(x, y)

	r := new(big.Float).SetFloat64(5)
	i := new(big.Float).SetFloat64(5)
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestSquare(t *testing.T) {
	x := fractal.NewComplex().SetFloat64(2, 5)
	square := fractal.NewComplex().Square(x)
	expected := fractal.NewComplex().Mul(x, x)

	assert.Equal(t, expected.Real, square.Real)
	assert.Equal(t, expected.Imaginary, square.Imaginary)
}
