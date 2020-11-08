package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mouckatron/buildservice/v2/internal/buildrunner"
)

var homeDir string
var buildsDir string

func init() {
	homeDir, _ = os.UserHomeDir()
	buildsDir = filepath.Join(homeDir, "builds")
}

func postBuild(c *gin.Context) {

	url := c.Query("url")
	name := c.DefaultQuery("name", "")
	branch := c.DefaultQuery("branch", "master")
	buildspec := c.DefaultQuery("buildspec", "buildspec.yml")

	// Settings
	settings := buildrunner.BuildSettings{
		Name:          name,
		URL:           url,
		Branch:        branch,
		BuildspecPath: buildspec,
	}

	// Setup
	err := buildrunner.Setup(&settings)

	// HTTP response based on settings/setup
	if err != nil {
		c.String(http.StatusInternalServerError, "Could not make build directory")
		return
	}

	c.String(http.StatusOK,
		fmt.Sprintf("http://localhost:8080/build/%s/%s",
			settings.Name,
			settings.ID))

	// Run
	go buildrunner.Run(settings)

	return
}

// getBuild returns a list of all builds available in the system
func getBuild(c *gin.Context) {
	file, err := os.Open(buildsDir)
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprint(err))
	}
	names, err := file.Readdirnames(0)
	if err != nil {
		c.String(http.StatusInternalServerError,
			fmt.Sprint(err))
	}
	c.String(http.StatusOK,
		strings.Join(names, "\n"))
}
