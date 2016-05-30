package fractal

import "fmt"

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
  return &Tile{tl, br, w, h, points}
}

func (t *Tile) getIndex(x, y int) int {
  return y * t.Width + x
}

func (t *Tile) getPoint(x, y int) *Point {
  index := t.getIndex(x, y)
  if nil == t.Points[index] {
    tl := t.TopLeft
    br := t.BottomRight
    dx := (br.Real - tl.Real) / float64(t.Width)
    dy := (br.Imaginary - tl.Imaginary) / float64(t.Height)
    r := tl.Real + float64(x)*dx
    i := tl.Imaginary + float64(y)*dy
    z := NewComplex()
    c := NewComplex64(r, i)
    t.Points[index] = NewPoint().Set(z, c)
  }
  return t.Points[index]
}

// Calculate calculates the points
func (t *Tile) Calculate(max uint64) *Tile {
  changed := true
  for changed {
    changed = false
    for y := 0; y < t.Height; y++ {
      for x := 0; x < t.Width; x++ {
        if t.skip(x, y, max) {
          continue
        }
        t.getPoint(x, y).Calculate(max)
        changed = true
      }
    }
  }
  return t
}

func (t *Tile) skip(x, y int, max uint64) bool {
  index := t.getIndex(x, y)
  if t.Points[index] != nil {
    return true
  }
  if x == 0 || x == t.Width - 1 || y == 0 || y == t.Height - 1 {
    return false
  }
  // Check if every neighbor has max iteration
  for i := -1; i < 2; i++ {
    for j := -1; j < 2; j++ {
      if i != 0 && j != 0 && !t.hasMaxItertion(x + i, y + j, max) {
        return false
      }
    }
  }
  return true
}

func (t *Tile) hasMaxItertion(x, y int, max uint64) bool {
  index := t.getIndex(x, y)
  p := t.Points[index]
  return nil == p || p.Iterations == max
}
// Print prints the tile
func (t *Tile) Print() {
  for index, p := range t.Points {
    if p != nil {
      fmt.Printf("%2d", p.Iterations)
    } else {
      fmt.Printf("xx")
    }
    fmt.Print(" ")
    if index%t.Width == t.Width-1 {
      fmt.Println("")
    }
  }
}
