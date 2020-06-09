package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRunCmd(t *testing.T) {
	t.Run("invalid arguments", func(t *testing.T) {
		env := make(Environment)
		cmd := make([]string, 0)

		require.Equal(t, ReturnCodeFail, RunCmd(cmd, env))
	})
}
