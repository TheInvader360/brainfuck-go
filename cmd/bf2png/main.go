package main

import (
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/TheInvader360/brainfuck-go/bf2png"
)

func main() {
	sourcePath := flag.String("source", "fixtures/text/helloworld.bf", "bf source path")
	outDir := flag.String("outdir", "transpiled/image/", "output directory path")
	flag.Parse()

	inFilename := filepath.Base(*sourcePath)
	outPathFilename := fmt.Sprintf("%s/%s.png", strings.TrimSuffix(*outDir, "/"), strings.TrimSuffix(inFilename, filepath.Ext(inFilename)))

	source, err := ioutil.ReadFile(*sourcePath)
	if err != nil {
		panic(err)
	}

	image := bf2png.NewTranspiler(string(source)).Transpile()

	file, err := os.Create(outPathFilename)
	if err != nil {
		panic(err)
	}
	png.Encode(file, image)
}
