package main

import (
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()

	//r.LoadHTMLFiles("./tmpl/home.html","./tmpl/userhome.html","./tmpl/scripts/home.js")
	//r.LoadHTMLGlob("static/**/*")

	r.LoadHTMLGlob("tmpl/*")
	r.Static("/static", "./static")
	r.Static("/video", "./video")

	r.Static("/statics/scripts", "./static/js")

	r.GET("/", homeHandler())

	r.POST("/", homeHandler())

	r.GET("/userhome", userHomeHandler())

	r.POST("/userhome", userHomeHandler())

	//
	r.POST("/api", apiHandler())

	r.GET("/videos/:vid-id", proxyVideoHandler())

	r.POST("/upload/:vid-id", proxyUploadHandler())

	r.Run(":8080")
}