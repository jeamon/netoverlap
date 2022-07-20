package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestExecute unit test the output result and exit code of logic behind the main function.
func TestExecute(t *testing.T) {
	// restore initial args after test.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	cases := []struct {
		name             string
		args             []string
		expectedExitCode int
		outputContains   string
	}{
		{
			"valid [version] argument",
			[]string{"version"},
			0,
			"Version:",
		},
		{
			"valid [--version] argument",
			[]string{"--version"},
			0,
			"Version:",
		},
		{
			"valid [-v] argument",
			[]string{"-v"},
			0,
			"Version:",
		},
		{
			"valid [help] argument",
			[]string{"help"},
			0,
			"Usage:",
		},
		{
			"valid [--help] argument",
			[]string{"--help"},
			0,
			"Usage:",
		},
		{
			"valid [-h] argument",
			[]string{"-h"},
			0,
			"Usage:",
		},
		{
			"valid prefixes arguments",
			[]string{"10.0.0.0/24", "10.0.0.0/8"},
			0,
			"superset",
		},
		{
			"invalid [-x] argument",
			[]string{"-x"},
			1,
			"Unknown command",
		},
		{
			"invalid syntax",
			[]string{"10.0.0.0/24", "10.0.0.0/8", "0.0.0.0/0"},
			1,
			"Invalid syntax",
		},
		{
			"invalid prefixes",
			[]string{"10.0.0.0/24", "10.0.0.0"},
			1,
			"failed",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			// make args parsing start at Args[1].
			os.Args = append([]string{tc.name}, tc.args...)
			buf := bytes.NewBuffer(nil)
			exitCode := Execute(buf)
			assert.Equal(t, tc.expectedExitCode, exitCode)
			assert.Contains(t, buf.String(), tc.outputContains)
		})
	}
}
