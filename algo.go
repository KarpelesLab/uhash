package main

import "hash"

var algoMap = make(map[string]*algo)

type algo struct {
	name    string
	desc    string
	factory func() hash.Hash
}

func reg(a *algo) {
	algoMap[a.name] = a
}
