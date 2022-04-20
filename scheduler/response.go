package main

import (
	"github.com/gin-gonic/gin"
)

func sendResponse(w gin.ResponseWriter, sc int, resp string) {
	w.WriteHeader(sc)
	//io.WriteString(w, resp)
}