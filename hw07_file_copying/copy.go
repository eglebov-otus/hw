package main

import (
	"errors"
	"io"
	"os"
	//"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath string, toPath string, offset, limit int64) error {
	src, err := os.Open(fromPath)

	if err != nil {
		panic(err)
	}

	info, err := os.Stat(fromPath)

	if err != nil {
		panic(err)
	}

	if 0 == info.Size() {
		return ErrUnsupportedFile
	}

	if offset > info.Size() {
		return ErrOffsetExceedsFileSize
	}

	dst, err := os.Create(toPath)

	if err != nil {
		panic(err)
	}

	_, err = src.Seek(offset, io.SeekStart)

	if err != nil {
		panic(err)
	}

	if limit > 0 {
		_, err = io.CopyN(dst, src, limit)

		if err != nil && err != io.EOF {
			panic(err)
		}
	} else {
		_, err = io.Copy(dst, src)

		if err != nil && err != io.EOF {
			panic(err)
		}
	}

	return nil
}
