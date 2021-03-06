package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const dirPath = "testdata/env"

func getEnv() Environment {
	env := make(Environment)
	env["BAR"] = EnvValue{"bar", false}
	env["EMPTY"] = EnvValue{"", false}
	env["FOO"] = EnvValue{"   foo\nwith new line", false}
	env["HELLO"] = EnvValue{"\"hello\"", false}
	env["UNSET"] = EnvValue{"", true}

	return env
}

func TestReadDir(t *testing.T) {
	env := getEnv()
	actual, err := ReadDir(dirPath)
	require.NoError(t, err)
	require.Equal(t, env, actual)
}

func TestInvalidDir(t *testing.T) {
	t.Run("empty path", func(t *testing.T) {
		_, err := ReadDir("")
		require.ErrorIs(t, err, ErrEmptyDir)
	})

	t.Run("invalid path", func(t *testing.T) {
		_, err := ReadDir("path")
		require.Error(t, err)
	})
}
