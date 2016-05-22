package fractal_test

import (
	"mandelbrot/fractal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPointToString(t *testing.T) {
	s := fractal.NewPoint()
	assert.Equal(t, "0+0i, i=0", s.String())
}

func TestIterate(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex().SetFloat64(0.7, 0.4)
	p := fractal.NewPoint().Set(z, c)

	p.Iterate()
	assert.Equal(t, "0.7+0.4i, i=1", p.String())
}

func TestCalculate(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex().SetFloat64(0.7, 0.4)
	p := fractal.NewPoint().Set(z, c)

	p.Calculate()
	assert.Equal(t, "0.8393+2.3776i, i=3", p.String())
}

func TestCalculate2(t *testing.T) {
	z := fractal.NewComplex()
	c := fractal.NewComplex().SetFloat64(-0.7, 0)
	p := fractal.NewPoint().Set(z, c)

	p.Calculate()
	assert.Equal(t, "-0.4746794345+0i, i=4096", p.String())
}
