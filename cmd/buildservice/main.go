package main

import (
	"io/ioutil"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/build/git", buildGit)

	router.Run()

}

func buildGit(c *gin.Context) {

	url := c.Query("url")

	dir, err := ioutil.TempDir("", "buildservice")
	if err != nil {
		log.Fatal(err)
		// gin return 500 server error here
	} else {
		log.Println(dir)
	}
	// defer os.RemoveAll(dir)

	log.Println(execCmd(dir, "git", "clone", url, dir))

	log.Println(execCmd(dir, "make", "build"))
}

func execCmd(dir string, cmd string, args ...string) (output string, err error) {
	c := exec.Command(cmd, args...)
	c.Dir = dir
	_output, err := c.Output()

	output = string(_output)

	return
}
