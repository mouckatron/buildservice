package buildrunner

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/kennygrant/sanitize"
	"github.com/mouckatron/go-buildspec/buildspec"
)

var homeDir string
var buildsDir string

func init() {
	homeDir, _ = os.UserHomeDir()
	buildsDir = filepath.Join(homeDir, "builds")
	os.Mkdir(buildsDir, os.FileMode(int(0750)))
}

// Setup gets the work directory and logger ready
func Setup(settings *BuildSettings) (err error) {

	// Name
	if settings.Name == "" {
		settings.Name = getNameFromURL(settings.URL)
	}

	// ID
	now := time.Now()
	settings.ID = now.Format("20060102-150405")

	// work directory
	settings.WorkingDir = filepath.Join(buildsDir, sanitize.Name(settings.Name), settings.ID)
	os.MkdirAll(settings.WorkingDir, os.FileMode(int(0750)))
	os.Mkdir(filepath.Join(settings.WorkingDir, "code"), os.FileMode(int(0750)))
	os.Mkdir(filepath.Join(settings.WorkingDir, "artifacts"), os.FileMode(int(0750)))

	// logger
	settings.LogFile, err = os.OpenFile(
		filepath.Join(settings.WorkingDir, "build.log"),
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	settings.Log = log.New(settings.LogFile, "", log.LUTC|log.Lmsgprefix|log.Ldate|log.Ltime)

	settings.Log.Println("Setup complete")

	return
}

// Run the main buildrunner function
func Run(settings BuildSettings) {

	defer settings.LogFile.Close()

	getCode(&settings)

	phases, err := getBuildCommands(&settings)

	settings.Log.Println(phases)
	settings.Log.Println(err)

	if err != nil {
		settings.Log.Fatal(err)
		return
	}

	for _, phase := range phases {
		settings.Log.Printf("Starting PHASE %s", phase.Name)

		for _, command := range phase.Commands {
			settings.Log.Printf("Executing command %s", command)
			output, err := execCmd(settings.CodeDir(), command)
			if err != nil {
				settings.Log.Fatal(err)
			}
			settings.Log.Println(output)
		}

		settings.Log.Printf("Finished PHASE %s", phase.Name)
	}

	settings.Log.Println("Build complete")

}

func getCode(settings *BuildSettings) (err error) {
	settings.Log.Println("getCode")
	if strings.Contains(settings.URL, "git") {
		getCodeWithGit(settings)
	}

	return
}

func getBuildCommands(settings *BuildSettings) (phases map[string]buildspec.Phase, err error) {

	settings.Log.Println("getBuildCommands")

	spec, err := buildspec.LoadFromFile(settings.BuildspecFile())
	if err != nil {
		settings.Log.Fatal(err)
		return
	}

	phases = spec.Phases

	return
}

func getCodeWithGit(settings *BuildSettings) (err error) {
	settings.Log.Printf("getCodeWithGit: %s", settings.URL)

	g, err := git.PlainClone(
		settings.CodeDir(),
		false,
		&git.CloneOptions{
			URL:          settings.URL,
			RemoteName:   settings.Branch,
			SingleBranch: true,
			Progress:     nil, //settings.Log.Writer(),
		})

	if err != nil {
		settings.Log.Println(err)
		return
	}

	ref, err := g.Head()
	if err != nil {
		settings.Log.Println(err)
		return
	}

	wt, err := g.Worktree()
	if err != nil {
		settings.Log.Println(err)
		return
	}

	err = wt.Checkout(&git.CheckoutOptions{
		Hash: ref.Hash(),
	})

	if err != nil {
		settings.Log.Println(err)
		return
	}
	return
}

func execCmd(dir string, cmd string, args ...string) (output string, err error) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	_output, err := c.CombinedOutput()

	output = string(_output)

	return
}
