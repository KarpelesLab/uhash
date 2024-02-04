package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash"
	"os"
	"strings"
)

type hashOutput interface {
	Append(filename string, hashAlgo *algo, value hash.Hash) error
	Finalize() error
}

func newOutput(style string) (hashOutput, error) {
	switch style {
	case "openssl":
		return opensslOutput{}, nil
	case "json":
		return &jsonOutput{data: make(map[string]map[string]any)}, nil
	default:
		return nil, fmt.Errorf("unsupported output format %s", style)
	}
}

// opensslOutput provides an output similar to openssl dgst (not exactly identical however)
type opensslOutput struct{}

func (opensslOutput) Append(filename string, hashAlgo *algo, value hash.Hash) error {
	if filename == "-" {
		// openssl writes "stdin" instead of -
		filename = "stdin"
	}
	var s string
	if ser, ok := value.(fmt.Stringer); ok {
		s = ser.String()
	} else {
		s = hex.EncodeToString(value.Sum(nil))
	}

	_, err := fmt.Fprintf(os.Stdout, "%s(%s)= %s\n", strings.ToUpper(hashAlgo.name), filename, s)
	return err
}

func (opensslOutput) Finalize() error {
	return nil
}

type jsonOutput struct {
	data map[string]map[string]any
}

func (j *jsonOutput) Append(filename string, hashAlgo *algo, value hash.Hash) error {
	if _, ok := j.data[filename]; !ok {
		j.data[filename] = make(map[string]any)
	}
	if obj, ok := value.(interface{ Value() any }); ok {
		j.data[filename][hashAlgo.name] = obj.Value()
	} else if ser, ok := value.(fmt.Stringer); ok {
		j.data[filename][hashAlgo.name] = ser.String()
	} else {
		j.data[filename][hashAlgo.name] = hex.EncodeToString(value.Sum(nil))
	}
	return nil
}

func (j *jsonOutput) Finalize() error {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	return enc.Encode(j.data)
}
