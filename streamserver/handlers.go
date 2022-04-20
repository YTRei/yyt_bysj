package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//func sendErrorResponse(w gin.ResponseWriter, sc int, errMsg string) {
//	//w.WriteHeader(sc)
//	io.WriteString(w, errMsg)
//}

type responseBodyWrite struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWrite) WriteString (s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func logResponseBody(c *gin.Context) {
	w := &responseBodyWrite{
		ResponseWriter: c.Writer,
		body:           &bytes.Buffer{},
	}
	c.Writer = w
	c.Next()

	fmt.Println("Response body: " + w.body.String())
}

func testPageHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "testpage.html", nil)
	}
}

func streamHandler() gin.HandlerFunc{
	return func(c *gin.Context) {
		//本地传输
	//	vid := c.Param("vid-id")
	//	vl := VIDEO_DIR + vid
	//
	//	video, err := os.Open(vl)
	//	if err != nil {
	//		c.Writer.WriteString("Error when try to open file")
	//		return
	//	}
	//
	//	c.Writer.Header().Set("Content-Type", "video/mp4")
	//	http.ServeContent(c.Writer, c.Request, "", time.Now(), video)
	//
	//	defer video.Close()

		// oss 传输
		log.Println("Entered the streamHandler")
		targetUrl := "http://yuan-videos.oss-cn-beijing.aliyuncs.com/videos/" + c.Params.ByName("vid-id")
		c.Redirect(http.StatusMovedPermanently, targetUrl)

	}
}

func uploadHandlers() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MAX_UPLOAD_SIZE)
		if err := c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
			c.Writer.WriteString("File is too big")
			return
		}

		file, _, err := c.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
			c.Writer.WriteString( "Internal Error")
			return
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Printf("Read file error: %v", err)
			c.Writer.WriteString( "Internal Error")
		}

		fn := c.Params.ByName("vid-id")
		err = ioutil.WriteFile(VIDEO_DIR + fn, data, 0666)
		if err != nil {
			log.Printf("Write file error: %v", err)
			c.Writer.WriteString("Internal Error")
			return
		}



		ossfn := "videos/" + fn
		path := "./videos/" + fn
		bn := "yuan-videos"
		ret := UploadToOss(ossfn, path, bn)
		if !ret {
			c.Writer.WriteString("Internal Error")
			return
		}

		os.Remove(path)

		c.Writer.WriteHeader(http.StatusCreated)
		c.Writer.WriteString("Upload successfuly")
	}
}