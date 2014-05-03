package fractals

import (
	"math"
	"math/rand"
	"testing"
)

func equilateralTransformer(line Line) []Line {
	third := 1.0 / 3.0
	slopeX := line.X2 - line.X1
	slopeY := line.Y2 - line.Y1

	line1 := NewLine(
		line.X1,
		line.Y1,
		line.X1 + third * slopeX,
		line.Y1 + third * slopeY,
	    line.Width)

	line4 := NewLine(
		line.X1 + (2 * third) * slopeX,
		line.Y1 + (2 * third) * slopeY,
		line.X2,
		line.Y2,
	    line.Width)

	x3, y3 := rotateLine(line1.X2, line1.Y2, line4.X1, line4.Y1, -60.0)

	line2 := NewLine(
		line1.X2,
		line1.Y2,
		line1.X2 + x3,
		line1.Y2 + y3,
	    line.Width)

	line3 := NewLine(
		line2.X2,
		line2.Y2,
		line4.X1,
		line4.Y1,
	    line.Width)

	return []Line{line1, line2, line3, line4}
}

func treeTransformer(line Line) []Line {
	ang1 := 15.0 + rand.Float64() * 10
	ang2 := -(15.0 + rand.Float64() * 10)
	len1 := 0.8 + rand.Float64() * 0.12
	len2 := 0.8 + rand.Float64() * 0.12
	x3, y3 := rotateLine(line.X1, line.Y1, line.X2, line.Y2, ang1)
	x4, y4 := rotateLine(line.X1, line.Y1, line.X2, line.Y2, ang2)
	line.Done = true
	return []Line{
		line,
		NewLine(
			line.X2,
			line.Y2,
			line.X2 + len1*x3,
			line.Y2 + len1*y3,
		    line.Width * 0.8),
		NewLine(
			line.X2,
			line.Y2,
			line.X2 + len2*x4,
			line.Y2 + len2*y4,
	        line.Width * 0.8)}
}

func rotateLine(x1, y1, x2, y2, deg float64) (float64, float64) {
	rad := math.Pi * deg / 180.0
	cos := math.Cos(rad)
	sin := math.Sin(rad)

	x3 := (x2 - x1) * cos - (y2 - y1) * sin
	y3 := (x2 - x1) * sin + (y2 - y1) * cos

	return x3, y3
}

func TestTree(t *testing.T) {
	initial := []Line{NewLine(0, 500, 0, 0, 70.0)}
	f := NewFractal(treeTransformer, initial)
	f.Evolve(10)
	f.ToFile("png/tree_fractal.png")
}

func TestEquilateral(t *testing.T) {
	initial := []Line{
		NewLine(0, 0, 500, 0, 2.0),
		NewLine(500, 0, 100, 50, 2.0),
		NewLine(100, 50, 0, 0, 2.0)}
	f := NewFractal(equilateralTransformer, initial)
	f.Evolve(8)
	f.ToFile("png/equal_fractal.png")
}
