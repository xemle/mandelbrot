package fractal

// Tile calculates a part of fractal
type Tile struct {
	TopLeft, BottomRight *Complex
	Width, Height        int
	Points               []*Point
}

// NewTile create new tile
func NewTile(tl, br *Complex, w, h int) *Tile {
	var len = w * h
	var points = make([]*Point, len, len)
	dx := (br.Real - tl.Real) / float64(w)
	dy := (br.Imaginary - tl.Imaginary) / float64(h)
	var index int
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			r := tl.Real + float64(x)*dx
			i := tl.Imaginary + float64(y)*dy
			z := NewComplex()
			c := NewComplex64(r, i)
			points[index] = NewPoint().Set(z, c)
			index++
		}
	}
	return &Tile{tl, br, w, h, points}
}

// Calculate calculates the points
func (t *Tile) Calculate(max uint64) *Tile {
	for _, p := range t.Points {
		p.Calculate(max)
	}

	return t
}
