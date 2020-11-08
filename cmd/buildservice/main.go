package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	log.SetFlags(0)
}

func main() {

	router := gin.Default()

	// create a new build
	router.POST("/build", postBuild)

	// get build history
	router.GET("/build", getBuild)

	// router.GET("/build/:name/", getBuildName)
	// router.GET("/build/:name/:id", getBuildNameId)
	// router.GET("/build/:name/:id/logs", getBuildLogs)
	// router.GET("/build/:name/:id/status", getBuildStatus)

	router.Run()

}
