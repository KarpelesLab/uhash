package main

import "github.com/jzelinskie/whirlpool"

func init() {
	reg(&algo{name: "whirlpool", desc: "ISO/IEC 10118-3:2004 Whirlpool", factory: whirlpool.New})
}
