package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// NAME of the application
const NAME = "buildservice"

func init() {

	home, _ := os.UserHomeDir()
	_ = os.Mkdir(filepath.Join(home, "artifacts"), os.FileMode(int(0755)))

}

func main() {

	router := gin.Default()

	router.POST("/build/git", buildGit)

	router.Run()

}

func buildGit(c *gin.Context) {

	url := c.Query("url")

	// get a directory and name the build id the same
	dir, id, err := workdirAndBuildID()

	if err != nil {
		c.String(http.StatusInternalServerError, "Could not make build directory")
		return
	}

	// return the id to the user
	c.String(http.StatusOK, fmt.Sprintf("http://localhost:8080/build/%s", id))
	// go routine to do the work
	go buildspecRunner(dir, id, url)

	return
}

func buildspecRunner(dir string, id string, url string) {
	defer os.RemoveAll(dir)
	log.Printf("Build %s starting\n", id)

	output, err := execCmd(dir, "git", "clone", url, dir)
	log.Println(output)
	log.Println(err)

	output, err = execCmd(dir, "make", "build")
	log.Println(output)
	log.Println(err)

	log.Printf("Build %s completed\n", id)
}

func workdirAndBuildID() (dir string, id string, err error) {
	dir, err = ioutil.TempDir("", NAME)
	if err != nil {
		log.Println(err)
		// gin return 500 server error here
	} else {
		log.Println(dir)
	}

	id = strings.Split(dir, NAME)[1]

	return
}

func execCmd(dir string, cmd string, args ...string) (output string, err error) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	_output, err := c.CombinedOutput()

	output = string(_output)

	return
}
