package fractal

import (
	"fmt"
	"math"
)

// Point of mandelbrot
type Point struct {
	Z, C       *Complex
	Iterations uint64
}

// NewPoint initialize new point
func NewPoint() *Point {
	z := NewComplex()
	c := NewComplex()
	var i uint64
	return &Point{z, c, i}
}

func (p *Point) String() string {
	return fmt.Sprintf("%s + %s, i=%d", p.Z, p.C, p.Iterations)
}

// Set sets the point
func (p *Point) Set(z, c *Complex) *Point {
	p.Z.Set(z)
	p.C.Set(c)

	return p
}

// Iterate calculate one iteration
func (p *Point) Iterate() {
	z2 := p.Z.Square(p.Z)
	p.Z = z2.Add(z2, p.C)
	p.Iterations++
}

// Calculate all iteration until break conditions are met
func (p *Point) Calculate(max uint64) *Point {
	for !p.Exceeds() && p.Iterations < max {
		p.Iterate()
	}
	return p
}

// Exceeds checks if the point exceeds
func (p *Point) Exceeds() bool {
	if math.Abs(p.Z.Real) < math.Sqrt2 && math.Abs(p.Z.Imaginary) < math.Sqrt2 {
		return false
	}
	if p.Z.Real >= 2 || p.Z.Imaginary > 2 {
		return true
	}

	r2 := p.Z.Real * p.Z.Real
	i2 := p.Z.Imaginary * p.Z.Imaginary
	d := r2 + i2
	result := 4 <= d
	return result
}
