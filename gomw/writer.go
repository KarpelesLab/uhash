package gomw

import (
	"io"
	"sync"
	"sync/atomic"
)

type multiWriter struct {
	w []io.Writer
}

func New[T io.Writer](w ...T) io.Writer {
	switch len(w) {
	case 0:
		return io.Discard
	case 1:
		return w[0]
	}

	writers := make([]io.Writer, len(w))
	for x, y := range w {
		writers[x] = y
	}
	return &multiWriter{w: writers}
}

// Write will simply write data to the different writers one by one in sequence.
// use ReadFrom to have something that will be more optimized
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

// ReadFrom will spawn one goroutine for each writer and perform some dark magic
// to parallelize writes to all the available writers.
//
// Ye who dareth approach this enigmatic conjuration, be warned:
// Herein lies complexity so vast and tangled, it doth threaten
// the very fabric of thy sanity. Proceed with caution, brave soul,
// lest ye wish to entangle thyself in its inexorable depths.
func (mw *multiWriter) ReadFrom(r io.Reader) (written int64, err error) {
	buf := make([]byte, 32*1024)
	var wbuf []byte // protected by s
	var s sync.RWMutex
	var wg sync.WaitGroup
	var errCnt uint32
	var errLk sync.Mutex
	var end bool // protected by s
	c := sync.NewCond(s.RLocker())
	s.Lock()
	defer s.Unlock()
	cnt := len(mw.w)

	for _, wr := range mw.w {
		go func(w io.Writer) {
			s.RLock()
			defer s.RUnlock()
			for {
				if len(wbuf) > 0 {
					wg.Done()
					n, er := w.Write(wbuf)
					if er == nil && n != len(wbuf) {
						er = io.ErrShortWrite
					}
					if er != nil {
						errLk.Lock()
						defer errLk.Unlock()
						err = er
						atomic.AddUint32(&errCnt, 1)
						return
					}
				}
				c.Wait()
				if end {
					return
				}
			}
		}(wr)
	}

	for {
		nr, er := r.Read(buf)
		if nr > 0 {
			wbuf = buf[:nr]
			wg.Add(cnt)
			c.Broadcast()
			s.Unlock()
			wg.Wait()
			s.Lock()

			if errCnt > 0 {
				end = true
				c.Broadcast()
				return
			}

			written += int64(nr)
		}
		if er != nil {
			end = true
			c.Broadcast()

			if er == io.EOF {
				return
			}
			err = er
			return
		}
	}
}
