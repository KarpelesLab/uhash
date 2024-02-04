package main

import "hash"

type concatHash struct {
	algos []hash.Hash
}

func newConcatFactory(factories ...func() hash.Hash) func() hash.Hash {
	return func() hash.Hash {
		return newConcat(factories...)
	}
}

func newConcat(factories ...func() hash.Hash) hash.Hash {
	h := &concatHash{}
	for _, f := range factories {
		h.algos = append(h.algos, f())
	}

	return h
}

func (c *concatHash) BlockSize() int {
	var res int
	for _, a := range c.algos {
		if n := a.BlockSize(); n > res {
			res = n
		}
	}
	return res
}

func (c *concatHash) Reset() {
	for _, a := range c.algos {
		a.Reset()
	}
}

func (c *concatHash) Sum(b []byte) []byte {
	for _, a := range c.algos {
		b = a.Sum(b)
	}
	return b
}

func (c *concatHash) Size() int {
	var res int
	for _, a := range c.algos {
		res += a.Size()
	}
	return res
}

func (c *concatHash) Write(b []byte) (int, error) {
	n := len(b)
	for _, a := range c.algos {
		if x, err := a.Write(b); err != nil {
			return x, err
		}
	}
	return n, nil
}
