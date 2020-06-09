package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
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

	n, err := calcBytesToCopy(fromPath, offset, limit)

	if err != nil {
		return err
	}

	dst, err := os.Create(toPath)

	if err != nil {
		panic(err)
	}

	_, err = src.Seek(offset, io.SeekStart)

	if err != nil {
		panic(err)
	}

	bar := pb.Full.Start64(n)
	barReader := bar.NewProxyReader(src)

	_, err = io.CopyN(dst, barReader, n)

	if err != nil && err != io.EOF {
		panic(err)
	}

	bar.Finish()

	return nil
}

func calcBytesToCopy(fromPath string, offset, limit int64) (int64, error) {
	info, err := os.Stat(fromPath)

	if err != nil {
		panic(err)
	}

	if info.Size() == 0 {
		return 0, ErrUnsupportedFile
	}

	if offset >= info.Size() {
		return 0, ErrOffsetExceedsFileSize
	}

	n := info.Size() - offset

	if limit > 0 && n > limit {
		n = limit
	}

	return n, nil
}
