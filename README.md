# uhash

Perform hashes of files, streams, etc in parallel.

    echo 'hello world' | sha256sum -

Becomes

    echo 'hello world' | uhash -with sha256 -

## Installing

    go install github.com/KarpelesLab/uhash@latest

## Supported hashing algorithms

```
$ ./uhash -list
List of supported hashing algorithms:
keccak256: Keccak-256 (legacy sha3)
keccak512: Keccak-512 (legacy sha3)
md4: MD4
md5: MD5
md5-sha1: MD5+SHA1 used for TLS RSA
ripemd160: RIPEMD160
sha1: SHA1
sha224: SHA224
sha256: SHA256
sha3-224: SHA3-224
sha3-256: SHA3-256
sha3-384: SHA3-384
sha3-512: SHA3-512
sha384: SHA384
sha512: SHA512
sha512-224: SHA512-224
sha512-256: SHA512-256
size: Number of bytes in stream
whirlpool: ISO/IEC 10118-3:2004 Whirlpool
```

## Testing

    openssl enc -pbkdf2 -aes-256-ctr -nosalt -pass pass:yourseed < /dev/zero 2>/dev/null | head -c $[1024*1024*1024] | ./uhash -with sha256,sha512 -

    openssl enc -pbkdf2 -aes-256-ctr -nosalt -pass pass:yourseed < /dev/zero 2>/dev/null | head -c $[1024*1024*1024] | sha256sum -b -
    openssl enc -pbkdf2 -aes-256-ctr -nosalt -pass pass:yourseed < /dev/zero 2>/dev/null | head -c $[1024*1024*1024] | sha512sum -b -

    sha256(stdin)=c5d53853949ac2faefe92b2478bd253b6380ffbf1dc6eb2b1eb5ac298f6cd7be
    sha512(stdin)=3a7c3a3fe7e60d2659c7fa16a3f319a61c781fa013f368cc6aaa6e415828d4eb6fd953f4bb0b0080c10f3474e8e97c2cf66527d852d8d9a4c0c5800e5242a5e3

