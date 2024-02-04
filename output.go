package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type hashOutput interface {
	Append(filename string, hashAlgo *algo, hashValue string) error
	Finalize() error
}

func newOutput(style string) (hashOutput, error) {
	switch style {
	case "openssl":
		return opensslOutput{}, nil
	case "json":
		return &jsonOutput{data: make(map[string]map[string]string)}, nil
	default:
		return nil, fmt.Errorf("unsupported output format %s", style)
	}
}

type opensslOutput struct{}

func (opensslOutput) Append(filename string, hashAlgo *algo, hashValue string) error {
	_, err := fmt.Fprintf(os.Stdout, "%s(%s): %s\n", hashAlgo.name, filename, hashValue)
	return err
}

func (opensslOutput) Finalize() error {
	return nil
}

type jsonOutput struct {
	data map[string]map[string]string
}

func (j *jsonOutput) Append(filename string, hashAlgo *algo, hashValue string) error {
	if _, ok := j.data[filename]; !ok {
		j.data[filename] = make(map[string]string)
	}
	j.data[filename][hashAlgo.name] = hashValue
	return nil
}

func (j *jsonOutput) Finalize() error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(j.data)
}
