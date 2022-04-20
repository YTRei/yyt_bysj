package main

import (
	"bysj_VEDIO/scheduler/dbops"
	"fmt"
	"github.com/gin-gonic/gin"
)

func vidDelRecHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		//var w gin.ResponseWriter
		//var p gin.Params

		//vid := p.ByName("vid-id")
		vid := c.Param("vid-id")

		fmt.Println("vid:" + vid)
		if len(vid) == 0 {
			//w.WriteHeader(400)
			return
		}

		err := dbops.AddVideoDeletionRecord(vid)
		if err != nil {
			//w.WriteHeader(500)
			//sendResponse(w, 500, "Internal server error")
			return
		}

		//w.WriteHeader(200)
		return
	}
}