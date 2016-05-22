package fractal

import (
	"fmt"
	"math"
	"math/big"
)

var sqrt2 = new(big.Float).SetFloat64(math.Sqrt2)

var square2 = big.NewFloat(4)
var defaultMaxIterations = big.NewInt(1 << 12)
var one = big.NewInt(1)

// Point of mandelbrot
type Point struct {
	Z, C       *Complex
	Iterations *big.Int
	Max        *big.Int
}

// NewPoint initialize new point
func NewPoint() *Point {
	z := NewComplex()
	c := NewComplex()
	i := new(big.Int)
	return &Point{z, c, i, defaultMaxIterations}
}

func (p *Point) String() string {
	return fmt.Sprintf("%s, i=%s", p.Z, p.Iterations)
}

// Set sets the point
func (p *Point) Set(z, c *Complex) *Point {
	p.Z.Set(z)
	p.C.Set(c)

	return p
}

// SetMaxIteration sets the maximum iteration
func (p *Point) SetMaxIteration(max *big.Int) *Point {
	p.Max = max

	return p
}

// SetPrec sets the precision of the numbers
func (p *Point) SetPrec(prec uint) *Point {
	p.Z = p.Z.SetPrec(prec)
	p.C = p.C.SetPrec(prec)

	return p
}

// Iterate calculate one iteration
func (p *Point) Iterate() {
	z2 := p.Z.Square(p.Z)
	p.Z = z2.Add(z2, p.C)
	p.Iterations = p.Iterations.Add(p.Iterations, one)
}

// Calculate all iteration until break conditions are met
func (p *Point) Calculate() *Point {
	for !p.Exceeds() && p.Iterations.Cmp(p.Max) < 0 {
		p.Iterate()
	}
	return p
}

// Exceeds checks if the point exceeds
func (p *Point) Exceeds() bool {
	if p.Z.Real.Cmp(sqrt2) < 0 && p.Z.Imaginary.Cmp(sqrt2) < 0 {
		return false
	}

	r2 := new(big.Float).Mul(p.Z.Real, p.Z.Real)
	i2 := new(big.Float).Mul(p.Z.Imaginary, p.Z.Imaginary)
	d := new(big.Float).Add(r2, i2)

	return d.Cmp(square2) >= 0
}
