package fractals

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"errors"
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"os"
)

type Line struct {
	X1, Y1, X2, Y2 float64
	Width          float64
	Done           bool
}

type TransformFunc func(Line) []Line

type Fractal struct {
	Transformer TransformFunc
	Lines       []Line
}

func NewFractal(transformer TransformFunc, initial []Line) Fractal {
	return Fractal{transformer, initial}
}

func (f *Fractal) Evolve(n int) {
	var nextLines, lastLines []Line
	lastLines = f.Lines
	for i := 0; i < n; i++ {
		nextLines = nil //make([]Line, 0, len(lastLines))
		for _, line := range lastLines {
			if line.Done {
				nextLines = append(nextLines, line)
			} else {
				nextLines = append(nextLines, f.Transformer(line)...)
			}
		}
		lastLines = nextLines
	}

	f.Lines = nextLines
}

func (f Fractal) ToFile(path string) error {
	if len(f.Lines) == 0 {
		return errors.New("Nothing to draw")
	}

	var minX, minY, maxX, maxY float64
	first := true

	for _, line := range f.Lines {
		if first {
			minX = math.Min(line.X1, line.X2)
			maxX = math.Max(line.X1, line.X2)
			minY = math.Min(line.Y1, line.Y2)
			maxY = math.Max(line.Y1, line.Y2)
			first = false
		} else {
			minX = math.Min(line.X1, math.Min(minX, line.X2))
			maxX = math.Max(line.X1, math.Max(maxX, line.X2))
			minY = math.Min(line.Y1, math.Min(minY, line.Y2))
			maxY = math.Max(line.Y1, math.Max(maxY, line.Y2))
		}
	}

	i := image.NewRGBA(image.Rect(0, 0, int(maxX-minX), int(maxY-minY)))
	gc := draw2d.NewGraphicContext(i)

	for _, line := range f.Lines {
		gc.SetLineWidth(float64(line.Width))
		gc.MoveTo(-minX+line.X1, -minY+line.Y1)
		gc.LineTo(-minX+line.X2, -minY+line.Y2)
		gc.Stroke()
	}

	saveToPngFile(path, i)
	return nil
}

func NewLine(x1, y1, x2, y2, width float64) Line {
	return Line{
		x1, y1, x2, y2, width, false}
}

func (line Line) Length() float64 {
	x, y := (line.X2 - line.X1), (line.Y2 - line.Y1)
	return math.Sqrt(x*x + y*y)
}

func (line Line) String() string {
	return fmt.Sprintf("%.3f %.3f -> %.3f %.3f (%.3f)",
		line.X1,
		line.Y1,
		line.X2,
		line.Y2,
		line.Length())
}

func saveToPngFile(filePath string, m image.Image) {
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
}
