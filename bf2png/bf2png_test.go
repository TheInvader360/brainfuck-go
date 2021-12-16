package bf2png

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestTranspile(t *testing.T) {
	type test struct {
		sourcePath        string
		expectedImagePath string
		tag               string
	}
	tests := []test{
		{sourcePath: "../fixtures/text-clean/ascii.bf", expectedImagePath: "../fixtures/image/ascii.png", tag: "ascii"},
		{sourcePath: "../fixtures/text-clean/beer.bf", expectedImagePath: "../fixtures/image/beer.png", tag: "beer"},
		{sourcePath: "../fixtures/text-clean/echo.bf", expectedImagePath: "../fixtures/image/echo.png", tag: "echo"},
		{sourcePath: "../fixtures/text-clean/helloworld.bf", expectedImagePath: "../fixtures/image/helloworld.png", tag: "helloworld"},
		{sourcePath: "../fixtures/text-clean/mandelbrot.bf", expectedImagePath: "../fixtures/image/mandelbrot.png", tag: "mandelbrot"},
		{sourcePath: "../fixtures/text-clean/rot13.bf", expectedImagePath: "../fixtures/image/rot13.png", tag: "rot13"},
		{sourcePath: "../fixtures/text-clean/squares.bf", expectedImagePath: "../fixtures/image/squares.png", tag: "squares"},
	}
	for _, tc := range tests {
		source, _ := ioutil.ReadFile(tc.sourcePath)
		expectedImageFile, _ := os.Open(tc.expectedImagePath)
		defer expectedImageFile.Close()
		expectedImage, _ := png.Decode(expectedImageFile)
		actualImage := NewTranspiler(string(source)).Transpile()
		bounds := actualImage.Bounds()
		for y := 0; y < bounds.Max.Y; y++ {
			for x := 0; x < bounds.Max.X; x++ {
				if actualImage.At(x, y) != expectedImage.At(x, y) {
					t.Errorf("expected:%v actual:%v (%s)", expectedImage, actualImage, tc.tag)
				}
			}
		}
	}
}
