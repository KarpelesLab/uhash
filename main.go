package main

import "flag"

var (
	hashAlgo = flag.String("with", "*", "choose which algorithms to use")
	listAlgo = flag.Bool("list", false, "list all available hash algorithms")
)

func main() {
	flag.Parse()
}
