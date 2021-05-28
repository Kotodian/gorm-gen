package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	dir    = flag.String("dir", "", "source dir")
	output = flag.String("output", "", "output dir")
)

func Usage() {
	_, _ = fmt.Fprintf(os.Stderr, "Usage of gorm-gen:\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()
	if len(*dir) == 0 {
		flag.Usage()
		os.Exit(2)
	}
	g := &Generator{output: *output}
	g.generate(*dir)
}
