package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"net/http"
	"os"
	"sort"
	"strings"
)

type runner struct {
	algos []*algo
	out   hashOutput
}

func newRunner(algos string, out hashOutput) (*runner, error) {
	r := &runner{out: out}

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
		return r.processReader(os.Stdin, fn)
	}

	if strings.HasPrefix(fn, "http://") || strings.HasPrefix(fn, "https://") {
		// get over http or https
		resp, err := http.Get(fn)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		if resp.StatusCode >= 300 {
			return fmt.Errorf("Invalid http status: %s", resp.Status)
		}

		return r.processReader(resp.Body, fn)
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
	mw := newMultiWriter(w...)
	_, err := io.Copy(mw, in)
	if err != nil {
		return err
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

		r.out.Append(inName, a, s)
	}
	return nil
}
