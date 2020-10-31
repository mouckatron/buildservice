package main

import (
	"os"
	"strings"
	"testing"
)

func TestExecCmdMultiple(t *testing.T) {
	var tests = []struct {
		dir, cmd string
		args     []string
		output   string
		error    bool
	}{
		{"nope", "nope", nil, "", true},
		{"/", "nope", nil, "", true},
		{"nope", "ls", nil, "", true},
		{"/", "echo", []string{"-n", "Hello", "World"}, "Hello World", false},
	}

	for _, tt := range tests {
		t.Run(tt.cmd, func(t *testing.T) {
			output, err := execCmd(tt.dir, tt.cmd, tt.args...)
			if output != tt.output {
				t.Errorf("got %s, want %s", output, tt.output)
			}
			if tt.error && err == nil {
				t.Errorf("got error %s", err)
			}
		})
	}
}

func TestWorkdirAndBuildID(t *testing.T) {

	prefix := "/tmp/buildservice"

	dir, id, err := workdirAndBuildID()
	defer os.RemoveAll(dir)

	if !strings.HasPrefix(dir, prefix) {
		t.Errorf("Temp directory does not start with %s, got %s", prefix, dir)
	}

	if !strings.HasSuffix(dir, id) || len(dir) == len(id) {
		t.Errorf("ID is not the temp directory suffix: wanted suffix of %s, got %s", dir, id)
	}

	if err != nil {
		t.Errorf("Got error creating temp directory: %s", err)
	}

}
