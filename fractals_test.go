package fractals

import (
	"testing"
)

func TestTree(t *testing.T) {
	initial := []Line{NewLine(0, 500, 0, 0, 70.0)}
	f := New(TreeTransformer, initial)
	f.Next(10)
	f.ToFile("png/tree_fractal.png")
}

func TestEquilateral(t *testing.T) {
	initial := []Line{
		NewLine(0, 0, 500, 0, 2.0),
		NewLine(500, 0, 100, 50, 2.0),
		NewLine(100, 50, 0, 0, 2.0)}
	f := New(EquilateralTransformer, initial)
	f.Next(8)
	f.ToFile("png/equal_fractal.png")
}

func TestGenExamples(t *testing.T) {
	initial := []Line{NewLine(0, 0, 100, 0, 2.0)}
	f := New(EquilateralTransformer, initial)
	f.ToFile("png/equal_ex1.png")
	f.Next(1)
	f.ToFile("png/equal_ex2.png")

	initial = []Line{NewLine(0, 100, 0, 0, 2.0)}
	f = New(TreeTransformer, initial)
	f.ToFile("png/tree_ex1.png")
	f.Next(1)
	f.ToFile("png/tree_ex2.png")
}
