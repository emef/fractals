fractals
========

fractal generator library in golang

**usage**

```go
import (
  "github.com/emef/fractals"
)

// create initial shape, just a triangle
lines := []fractals.Line{
		fractals.NewLine(0, 0, 500, 0, 2.0),
		fractals.NewLine(500, 0, 100, 50, 2.0),
		fractals.NewLine(100, 50, 0, 0, 2.0)}
		
// fractal object, will transform each line according to EquilateralTransformer in each step
f := fractals.New(fractals.EquilateralTransformer, lines)

// go through 8 evolutions
f.Evolve(8)

// save fractal to png
f.ToFile("png/equal_fractal.png")
```
