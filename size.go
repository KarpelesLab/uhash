package main

import (
	"encoding/binary"
	"hash"
	"strconv"
)

func init() {
	reg(&algo{name: "size", desc: "Number of bytes in stream", factory: newSizeWriter})
}

func newSizeWriter() hash.Hash {
	return &sizeWriter{}
}

// sizeWriter is a simple method that acts like a hash.Hash but only measures size
type sizeWriter struct {
	bytes uint64
}

func (sw *sizeWriter) BlockSize() int {
	return 1
}

func (sw *sizeWriter) Reset() {
	sw.bytes = 0
}

func (sw *sizeWriter) Size() int {
	return 8
}

func (sw *sizeWriter) Sum(b []byte) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], sw.bytes)
	return append(b, buf[:]...)
}

func (sw *sizeWriter) Write(b []byte) (int, error) {
	n := len(b)
	sw.bytes += uint64(n)
	return n, nil
}

// String implements io.Stringer so we display the size in base 10 instead of a binary value
func (sw *sizeWriter) String() string {
	return strconv.FormatUint(sw.bytes, 10)
}
