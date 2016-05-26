package fractal_test

import (
	"mandelbrot/fractal"
	"testing"

	"github.com/stretchr/testify/assert"
)

var max uint64 = 1 << 12

func TestPointToString(t *testing.T) {
	s := fractal.NewPoint()
	assert.Equal(t, "0.000000+0.000000i + 0.000000+0.000000i, i=0", s.String())
}

func TestPointExceeds(t *testing.T) {
	z := fractal.NewComplex64(2.0, -2.0)
	c := fractal.NewComplex()
	p := fractal.NewPoint().Set(z, c)

	assert.Equal(t, true, p.Exceeds())
}

func TestPointIterate(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex64(0.7, 0.4)
	p := fractal.NewPoint().Set(z, c)

	p.Iterate()
	assert.Equal(t, "0.700000+0.400000i + 0.700000+0.400000i, i=1", p.String())
}

func TestPointCalculate(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex64(0.7, 0.4)
	p := fractal.NewPoint().Set(z, c)

	p.Calculate(max)
	assert.Equal(t, "0.839300+2.377600i + 0.700000+0.400000i, i=3", p.String())
}

func TestPointCalculate2(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex64(-0.7, 0)
	p := fractal.NewPoint().Set(z, c)

	p.Calculate(max)
	assert.Equal(t, "-0.474679+0.000000i + -0.700000+0.000000i, i=4096", p.String())
}
