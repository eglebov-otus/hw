package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"
)

var ErrInvalidDir = errors.New("failed to read directory")
var ErrEmptyDir = errors.New("empty dir")
var ErrInvalidFile = errors.New("failed to read file")

type Environment map[string]string

// ReadDir reads a specified directory and returns map of env variables.
// Variables represented as files where filename is name of variable, file first line is a value.
func ReadDir(dir string) (Environment, error) {
	files, err := ioutil.ReadDir(dir)

	res := make(Environment)

	if err != nil {
		return res, ErrInvalidDir
	}

	if len(files) == 0 {
		return res, ErrEmptyDir
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))

		if err != nil {
			return res, ErrInvalidFile
		}

		lines := bytes.Split(data, []byte("\n"))

		var value []byte

		if len(lines) > 0 {
			value = lines[0]
			value = bytes.Replace(value, []byte("\x00"), []byte("\n"), -1)
			value = bytes.TrimRight(value, " ")
			value = bytes.TrimRight(value, "\t")
		}

		res[file.Name()] = string(value)
	}

	return res, nil
}
