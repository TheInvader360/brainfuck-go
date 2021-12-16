package png2bf

import (
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestTranspile(t *testing.T) {
	type test struct {
		imagePath          string
		expectedSourcePath string
		tag                string
	}
	tests := []test{
		{imagePath: "../fixtures/image/ascii.png", expectedSourcePath: "../fixtures/text-clean/ascii.bf", tag: "ascii"},
		{imagePath: "../fixtures/image/beer.png", expectedSourcePath: "../fixtures/text-clean/beer.bf", tag: "beer"},
		{imagePath: "../fixtures/image/echo.png", expectedSourcePath: "../fixtures/text-clean/echo.bf", tag: "echo"},
		{imagePath: "../fixtures/image/helloworld.png", expectedSourcePath: "../fixtures/text-clean/helloworld.bf", tag: "helloworld"},
		{imagePath: "../fixtures/image/mandelbrot.png", expectedSourcePath: "../fixtures/text-clean/mandelbrot.bf", tag: "mandelbrot"},
		{imagePath: "../fixtures/image/rot13.png", expectedSourcePath: "../fixtures/text-clean/rot13.bf", tag: "rot13"},
		{imagePath: "../fixtures/image/squares.png", expectedSourcePath: "../fixtures/text-clean/squares.bf", tag: "squares"},
	}
	for _, tc := range tests {
		imageFile, _ := os.Open(tc.imagePath)
		defer imageFile.Close()
		image, _ := png.Decode(imageFile)
		expectedSource, _ := ioutil.ReadFile(tc.expectedSourcePath)
		actualSource := NewTranspiler(image).Transpile()
		if actualSource != string(expectedSource) {
			t.Errorf("expected:%s actual:%s (%s)", expectedSource, actualSource, tc.tag)
		}
	}
}
