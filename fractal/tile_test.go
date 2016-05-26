package fractal_test

import (
	"fmt"
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

	for index, p := range tile.Points {
		fmt.Printf("%2d", p.Iterations)
		fmt.Print(" ")
		if index%s == s-1 {
			fmt.Println("")
		}
	}
}
