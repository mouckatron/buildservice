package buildrunner

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func cleanupAfterSetup(b *BuildSettings) {
	b.LogFile.Close()
	os.RemoveAll("/tmp/buildservice")
}

func TestSetupName(t *testing.T) {
	b := BuildSettings{URL: "github.com/mouckatron/buildservice.git"}

	expected := "buildservice"

	homeDir = "/tmp"
	buildsDir = "/tmp"
	Setup(&b)

	if b.Name != expected {
		t.Errorf("Bad Name: wanted %s, got %s", expected, b.Name)
	}

	cleanupAfterSetup(&b)
}

// This seems a bit pointless...but in the name of coverage
func TestSetupID(t *testing.T) {
	b := BuildSettings{}

	now := time.Now()
	expected := now.Format("20060102-150405")

	homeDir = "/tmp"
	buildsDir = "/tmp"
	Setup(&b)

	if b.ID != expected {
		t.Errorf("Bad ID: wanted %s, got %s (this may also be a timing issue if it rolled over the second)",
			expected, b.ID)
	}

	cleanupAfterSetup(&b)
}

func TestSetupWorkDirectory(t *testing.T) {
	b := BuildSettings{
		Name: "buildservice"}

	wdExpected := "/tmp/buildservice"

	homeDir = "/tmp"
	buildsDir = "/tmp"
	Setup(&b)

	if !strings.HasPrefix(b.WorkingDir, wdExpected) {
		t.Errorf("Working directory bad: expected %s, got %s", wdExpected, b.WorkingDir)
	}

	cleanupAfterSetup(&b)
}

func TestSetupLogging(t *testing.T) {
	logFilename := "build.log"
	b := BuildSettings{
		Name: "buildservice"}

	homeDir = "/tmp"
	buildsDir = "/tmp"

	Setup(&b)

	if _, err := os.Stat(filepath.Join(b.WorkingDir, logFilename)); err != nil {
		t.Errorf("%s did not exist", logFilename)
	}

	cleanupAfterSetup(&b)
}

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
