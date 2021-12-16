package png2bf

import (
	"image"
	"image/color"
)

type Transpiler struct {
	image image.Image // source image
}

func NewTranspiler(image image.Image) *Transpiler {
	return &Transpiler{
		image: image,
	}
}

func (t *Transpiler) Transpile() string {
	source := ""

	bounds := t.image.Bounds()
	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			switch t.image.At(x, y) {
			case color.RGBA{0, 255, 0, 255}: // Green
				source += ">"
			case color.RGBA{255, 255, 0, 255}: // Yellow
				source += "<"
			case color.RGBA{255, 0, 0, 255}: // Red
				source += "+"
			case color.RGBA{0, 0, 255, 255}: // Blue
				source += "-"
			case color.RGBA{255, 0, 255, 255}: // Magenta
				source += "."
			case color.RGBA{0, 255, 255, 255}: // Cyan
				source += ","
			case color.RGBA{0, 0, 0, 255}: // Black
				source += "["
			case color.RGBA{255, 255, 255, 255}: // White
				source += "]"
			}
		}
	}

	return source
}
