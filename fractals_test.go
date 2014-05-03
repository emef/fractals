package fractals

import (
	"testing"
)

func TestTree(t *testing.T) {
	initial := []Line{NewLine(0, 500, 0, 0, 70.0)}
	f := NewFractal(TreeTransformer, initial)
	f.Next(10)
	f.ToFile("png/tree_fractal.png")
}

func TestEquilateral(t *testing.T) {
	initial := []Line{
		NewLine(0, 0, 500, 0, 2.0),
		NewLine(500, 0, 100, 50, 2.0),
		NewLine(100, 50, 0, 0, 2.0)}
	f := NewFractal(EquilateralTransformer, initial)
	f.Next(8)
	f.ToFile("png/equal_fractal.png")
}
