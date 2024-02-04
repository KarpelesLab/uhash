package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	hashAlgo = flag.String("with", "*", "choose which algorithms to use")
	listAlgo = flag.Bool("list", false, "list all available hash algorithms")
)

func main() {
	flag.Parse()

	p, err := newRunner(*hashAlgo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %s\n", err)
		os.Exit(1)
	}

	for _, fn := range flag.Args() {
		// process fn
		p.process(fn)
	}
}
