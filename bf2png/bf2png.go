package bf2png

import (
	"image"
	"image/color"
	"math"
)

type Transpiler struct {
	source string // bf source (e.g. "--[+++++++<---->>-->+>+>+<<<<]<.>++++[-<++++<++>>>->--<<]>>-.>--..>+.<<<.<<-.>>+>->>.+++[.<]<<++.")
}

func NewTranspiler(source string) *Transpiler {
	return &Transpiler{
		source: source,
	}
}

func (t *Transpiler) Transpile() *image.RGBA {
	colors := []color.RGBA{}
	for _, c := range t.source {
		switch c {
		case '>':
			colors = append(colors, color.RGBA{0, 255, 0, 255}) // Green
		case '<':
			colors = append(colors, color.RGBA{255, 255, 0, 255}) // Yellow
		case '+':
			colors = append(colors, color.RGBA{255, 0, 0, 255}) // Red
		case '-':
			colors = append(colors, color.RGBA{0, 0, 255, 255}) // Blue
		case '.':
			colors = append(colors, color.RGBA{255, 0, 255, 255}) // Magenta
		case ',':
			colors = append(colors, color.RGBA{0, 255, 255, 255}) // Cyan
		case '[':
			colors = append(colors, color.RGBA{0, 0, 0, 255}) // Black
		case ']':
			colors = append(colors, color.RGBA{255, 255, 255, 255}) // White
		}
	}

	width := int(math.Ceil(math.Sqrt(float64(len(colors)))))
	height := int(math.Ceil(float64(len(colors)) / float64(width)))
	image := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			index := y*width + x
			if index < len(colors) {
				image.Set(x, y, colors[index])
			} else {
				image.Set(x, y, color.RGBA{128, 128, 128, 255})
			}
		}
	}
	// fmt.Printf("raw source length = %d, stripped source length = %d, total pixels = %d\n", len(t.source), len(colors), width*height)

	return image
}
