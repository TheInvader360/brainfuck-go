package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/TheInvader360/brainfuck-go/vm"
)

func main() {
	sourcePath := flag.String("source", "programs/text/helloworld.bf", "bf source path")
	debug := flag.Bool("debug", false, "enable debug output")
	flag.Parse()

	source, err := ioutil.ReadFile(*sourcePath)
	if err != nil {
		panic(err)
	}

	started := time.Now()
	vm.NewVM(string(source), os.Stdin, os.Stdout, *debug).Run()
	fmt.Println("\nCompleted in", time.Since(started))
}
