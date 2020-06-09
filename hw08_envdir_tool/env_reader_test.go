package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestReadDir(t *testing.T) {
	t.Run("invalid dir", func(t *testing.T) {
		_, err := ReadDir("/invalid_dir")

		require.Equal(t, ErrInvalidDir, err)
	})

	t.Run("empty dir", func(t *testing.T) {
		_ = os.Mkdir("testdata/empty_dir", os.ModePerm)
		defer func() {
			_ = os.Remove("testdata/empty_dir")
		}()

		_, err := ReadDir("testdata/empty_dir")

		require.Equal(t, ErrEmptyDir, err)
	})

	t.Run("invalid file", func(t *testing.T) {
		_, _ = os.Create("testdata/invalid_file")
		_ = os.Chmod("testdata/invalid_file", 0000)
		defer func() {
			_ = os.Remove("testdata/invalid_file")
		}()

		_, err := ReadDir("testdata")

		require.Equal(t, ErrInvalidFile, err)
	})
}
