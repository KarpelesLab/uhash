package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
	"golang.org/x/crypto/md4"
	"golang.org/x/crypto/ripemd160"
	"golang.org/x/crypto/sha3"
)

func init() {
	reg(&algo{name: "md4", desc: "MD4", factory: md4.New})
	reg(&algo{name: "md5", desc: "MD5", factory: md5.New})
	reg(&algo{name: "sha1", desc: "SHA1", factory: sha1.New})
	reg(&algo{name: "md5-sha1", desc: "MD5+SHA1 used for TLS RSA", factory: newConcatFactory(md5.New, sha1.New)})
	reg(&algo{name: "sha256", desc: "SHA256", factory: sha256.New})
	reg(&algo{name: "sha224", desc: "SHA224", factory: sha256.New224})
	reg(&algo{name: "sha384", desc: "SHA384", factory: sha512.New384})
	reg(&algo{name: "sha512", desc: "SHA512", factory: sha512.New})
	reg(&algo{name: "sha512-224", desc: "SHA512-224", factory: sha512.New512_224})
	reg(&algo{name: "sha512-256", desc: "SHA512-256", factory: sha512.New512_256})
	reg(&algo{name: "ripemd-160", alias: []string{"ripemd", "ripemd160", "rmd160"}, desc: "RIPE Message Digest", factory: ripemd160.New})
	reg(&algo{name: "sha3-224", desc: "SHA3-224", factory: sha3.New224})
	reg(&algo{name: "sha3-256", desc: "SHA3-256", factory: sha3.New256})
	reg(&algo{name: "sha3-384", desc: "SHA3-384", factory: sha3.New384})
	reg(&algo{name: "sha3-512", desc: "SHA3-512", factory: sha3.New512})
	reg(&algo{name: "shake128", desc: "SHA3 SHAKE128", factory: func() hash.Hash { return sha3.NewShake128() }})
	reg(&algo{name: "shake256", desc: "SHA3 SHAKE256", factory: func() hash.Hash { return sha3.NewShake256() }})
	reg(&algo{name: "keccak256", desc: "Keccak-256 (legacy sha3)", factory: sha3.NewLegacyKeccak256})
	reg(&algo{name: "keccak512", desc: "Keccak-512 (legacy sha3)", factory: sha3.NewLegacyKeccak512})
	reg(&algo{name: "blake2s-256", alias: []string{"blake2s256"}, desc: "BLAKE2s-256", factory: mustBlake(blake2s.New256)})
	reg(&algo{name: "blake2b-512", alias: []string{"blake2b512"}, desc: "BLAKE2b-512", factory: mustBlake(blake2b.New512)})
}

func mustBlake(f func([]byte) (hash.Hash, error)) func() hash.Hash {
	return func() hash.Hash {
		v, err := f(nil)
		if err != nil {
			panic(err)
		}
		return v
	}
}
