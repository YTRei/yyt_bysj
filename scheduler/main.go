package main

import (
	"bysj_VEDIO/scheduler/taskrunner"
	"github.com/gin-gonic/gin"
)

func main() {
	go taskrunner.Start()

	r := gin.Default()

	r.GET("/video-delete-record/:vid-id", vidDelRecHandler())

	r.Run(":9001")
}