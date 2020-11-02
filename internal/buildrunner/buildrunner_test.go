package buildrunner

import "testing"

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
