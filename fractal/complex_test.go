package fractal_test

import (
	"mandelbrot/fractal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexToString(t *testing.T) {
	s := fractal.NewComplex()
	assert.Equal(t, "0.000000+0.000000i", s.String())
}

func TestNewComplex64(t *testing.T) {
	z := fractal.NewComplex64(1, 0)
	r := 1.0
	i := 0.0
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestAdd(t *testing.T) {
	x := fractal.NewComplex64(1, 0)
	y := fractal.NewComplex64(0, 2)
	z := fractal.NewComplex().Add(x, y)

	r := 1.0
	i := 2.0
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestMul(t *testing.T) {
	x := fractal.NewComplex64(1, 2)
	y := fractal.NewComplex64(3, -1)
	z := fractal.NewComplex().Mul(x, y)

	r := 5.0
	i := 5.0
	assert.Equal(t, r, z.Real)
	assert.Equal(t, i, z.Imaginary)
}

func TestSquare(t *testing.T) {
	x := fractal.NewComplex64(2, 5)
	square := fractal.NewComplex().Square(x)
	expected := fractal.NewComplex().Mul(x, x)

	assert.Equal(t, expected.Real, square.Real)
	assert.Equal(t, expected.Imaginary, square.Imaginary)
}
