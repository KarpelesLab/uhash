package main

import (
	"crypto/sha256"
	"crypto/sha512"
)

func init() {
	reg(&algo{name: "sha256", desc: "SHA256", factory: sha256.New})
	reg(&algo{name: "sha224", desc: "SHA224", factory: sha256.New224})
	reg(&algo{name: "sha384", desc: "SHA384", factory: sha512.New384})
	reg(&algo{name: "sha512", desc: "SHA512", factory: sha512.New})
}
