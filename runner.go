package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"sort"
	"strings"
)

type runner struct {
	algos []*algo
}

func newRunner(algos string) (*runner, error) {
	r := &runner{}

	if algos == "*" {
		var v []string
		for k := range algoMap {
			v = append(v, k)
		}
		sort.Strings(v)
		for _, k := range v {
			r.algos = append(r.algos, algoMap[k])
		}
		return r, nil
	}

	a := strings.Split(algos, ",")
	for _, k := range a {
		if obj, ok := algoMap[k]; ok {
			r.algos = append(r.algos, obj)
		} else {
			return nil, fmt.Errorf("unknown hashing algorithm: %s", k)
		}
	}

	return r, nil
}

func (r *runner) process(fn string) error {
	if fn == "-" {
		return r.processReader(os.Stdin, "stdin")
	}
	f, err := os.Open(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	return r.processReader(f, fn)
}

func (r *runner) processReader(in io.Reader, inName string) error {
	var w []hash.Hash
	for _, a := range r.algos {
		w = append(w, a.factory())
	}

	buf := make([]byte, 8192)

	for {
		n, err := in.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		for _, out := range w {
			_, err = out.Write(buf[:n])
			if err != nil {
				return err
			}
		}
	}

	// finished
	for n, a := range r.algos {
		var s string
		out := w[n]
		if ser, ok := out.(fmt.Stringer); ok {
			s = ser.String()
		} else {
			s = hex.EncodeToString(out.Sum(nil))
		}

		fmt.Fprintf(os.Stdout, "%s(%s)=%s\n", a.name, inName, s)
	}
	return nil
}
