package server

import (
	"fmt"
	"mandelbrot/fractal"
)

// TileCommand holds tile information to render
type TileCommand struct {
	Top          float64 `json:"top"`
	Left         float64 `json:"left"`
	Bottom       float64 `json:"bottom"`
	Right        float64 `json:"right"`
	Width        int     `json:"width"`
	Height       int     `json:"height"`
	MaxIteration uint64  `json:"maxIteration"`
}

// TileCommandResult holds the result
type TileCommandResult struct {
	TileCommand
	Iterations []uint64 `json:"iterations"`
}

// Calculate calculate tile
func (t *TileCommand) Calculate() *TileCommandResult {
	tl := fractal.NewComplex64(t.Left, t.Top)
	br := fractal.NewComplex64(t.Right, t.Bottom)

	fmt.Println(tl)
	fmt.Println(br)
	tile := fractal.NewTile(tl, br, t.Width, t.Height)
	tile.Calculate(t.MaxIteration)

	var iterations = make([]uint64, len(tile.Points), len(tile.Points))
	for index, p := range tile.Points {
		iterations[index] = p.Iterations
	}
	//tile.Print()
	return &TileCommandResult{TileCommand: *t, Iterations: iterations}
}
