package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var (
	hashAlgo     = flag.String("with", "*", "choose which algorithms to use")
	listAlgo     = flag.Bool("list", false, "list all available hash algorithms")
	outputFormat = flag.String("format", "openssl", "select output format for result")
)

func main() {
	flag.Parse()

	if *listAlgo {
		// list algos
		fmt.Fprintf(os.Stderr, "List of supported hashing algorithms:\n")
		var v []string
		for k := range algoMap {
			v = append(v, k)
		}
		sort.Strings(v)
		for _, k := range v {
			a := algoMap[k]
			fmt.Fprintf(os.Stderr, "%s: %s\n", a.name, a.desc)
		}
		return
	}

	out, err := newOutput(*outputFormat)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %s\n", err)
		os.Exit(1)
	}
	defer out.Finalize()

	p, err := newRunner(*hashAlgo, out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to initialize: %s\n", err)
		os.Exit(1)
	}

	files := flag.Args()
	if len(files) == 0 {
		// read from stdin by default
		files = []string{"-"}
	}

	for _, fn := range files {
		// process fn
		p.process(fn)
	}
}
