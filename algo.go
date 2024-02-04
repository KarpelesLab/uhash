package main

import "hash"

var (
	algoMap  = make(map[string]*algo)
	aliasMap = make(map[string]*algo)
)

type algo struct {
	name    string
	alias   []string
	desc    string
	factory func() hash.Hash
}

func reg(a *algo) {
	algoMap[a.name] = a
	for _, alias := range a.alias {
		aliasMap[alias] = a
	}
}
