package buildrunner

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// BuildSettings to pass to the work functions
type BuildSettings struct {
	Name          string // Build group
	ID            string // Timestamp
	URL           string
	Branch        string
	BuildspecPath string
	WorkingDir    string // full path including Name and ID
	Log           *log.Logger
	LogFile       *os.File
}

// CodeDir returns the code directory from BuildSettings
func (b *BuildSettings) CodeDir() string {
	return filepath.Join(b.WorkingDir, "code")
}

// ArtifactDir returns the artifact directory from BuildSettings
func (b *BuildSettings) ArtifactDir() string {
	return filepath.Join(b.WorkingDir, "artifacts")
}

// BuildspecFile returns the buildspec file path from BuildSettings
func (b *BuildSettings) BuildspecFile() string {
	return filepath.Join(b.CodeDir(), b.BuildspecPath)
}

// ToString returns a debug log printable version of the struct
func (b *BuildSettings) ToString() string {
	return fmt.Sprintf("Name: %s, ID: %s, URL: %s, Branch: %s, BuildspecPath: %s, WorkingDir: %s",
		b.Name,
		b.ID,
		b.URL,
		b.Branch,
		b.BuildspecPath,
		b.WorkingDir)
}
