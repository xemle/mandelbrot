package fractal_test

import (
	"mandelbrot/fractal"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTile(t *testing.T) {
	tl := fractal.NewComplex64(-2.0, 2.0)
	br := fractal.NewComplex64(2.0, -2.0)
	tile := fractal.NewTile(tl, br, 4, 4)

	assert.Equal(t, 16, len(tile.Points))
}

func TestTileCalculate2(t *testing.T) {
	tl := fractal.NewComplex64(-2.0, 1.5)
	br := fractal.NewComplex64(1.0, -1.5)
	s := 25
	tile := fractal.NewTile(tl, br, s, s)
	var max uint64 = 99

	tile.Calculate(max)

	tile.Print()
}
