package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCopy(t *testing.T) {
	t.Run("offset exceeds file size", func(t *testing.T) {
		require.Equal(t, ErrOffsetExceedsFileSize, Copy("testdata/input.txt", "/tmp/out.txt", 1000000, 0))
	})

	t.Run("unsupported file", func(t *testing.T) {
		require.Equal(t, ErrUnsupportedFile, Copy("/dev/urandom", "/tmp/out.txt", 1000000, 0))
	})
}
