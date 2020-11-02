package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mouckatron/buildservice/v2/internal/buildrunner"
)

func main() {

	router := gin.Default()

	router.POST("/build/", build)

	router.Run()

}

func build(c *gin.Context) {

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
