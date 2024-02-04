package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

var (
	hashAlgo = flag.String("with", "*", "choose which algorithms to use")
	listAlgo = flag.Bool("list", false, "list all available hash algorithms")
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
