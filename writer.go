package main

import "io"

type multiWriter struct {
	w []io.Writer
}

func newMultiWriter[T io.Writer](w ...T) io.Writer {
	writers := make([]io.Writer, len(w))
	for x, y := range w {
		writers[x] = y
	}
	return &multiWriter{w: writers}
}

func (mw *multiWriter) Write(b []byte) (int, error) {
	l := len(b)
	for _, w := range mw.w {
		n, err := w.Write(b)
		if err != nil {
			return n, err
		}
		if n != l {
			return n, io.ErrShortWrite
		}
	}
	return l, nil
}
