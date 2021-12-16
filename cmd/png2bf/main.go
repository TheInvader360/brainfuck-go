package main

import (
	"errors"
	"flag"
	"fmt"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/TheInvader360/brainfuck-go/png2bf"
)

func main() {
	imagePath := flag.String("source", "fixtures/image/helloworld.png", "source image path")
	outDir := flag.String("outdir", "transpiled/text/", "output directory path")
	flag.Parse()

	inFilename := filepath.Base(*imagePath)
	outPathFilename := fmt.Sprintf("%s/%s.bf", strings.TrimSuffix(*outDir, "/"), strings.TrimSuffix(inFilename, filepath.Ext(inFilename)))

	inFile, err := os.Open(*imagePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	image, err := png.Decode(inFile)
	if err != nil {
		panic(err)
	}

	source := png2bf.NewTranspiler(image).Transpile()

	if len(source) == 0 {
		panic(errors.New("zero length output"))
	}
	err = ioutil.WriteFile(outPathFilename, []byte(source), 0777)
	if err != nil {
		panic(err)
	}
}
