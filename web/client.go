package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"yyt/config"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, c *gin.Context, w gin.ResponseWriter) {
	var resp *http.Response
	var err error

	u, _ := url.Parse(b.Url)
	u.Host = config.GetLbAddr() + ":" + u.Port()
	newUrl := u.String()

	//fmt.Println(b.Method,"++++",b.Url)
	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("GET", newUrl, nil)
		//fmt.Println("1111111")
		req.Header = c.Request.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("%v",err)
			return
		}
		normalResponse(w, resp)

	case http.MethodPost:
		req, _ := http.NewRequest("POST", newUrl, bytes.NewBuffer([]byte(b.ReqBody)))

		req.Header = c.Request.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("%v",err)
			return
		}
		//fmt.Println("22222222")
		normalResponse(w, resp)

	case http.MethodDelete:
		req, _ := http.NewRequest("Delete", newUrl, nil)
		req.Header = c.Request.Header
		resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("%v",err)
			return
		}
		normalResponse(w, resp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.WriteString("Bad api request")
		return
	}
}

func normalResponse(w gin.ResponseWriter,r *http.Response) {
	res, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		w.WriteString("internal service error")
		return
	}

	w.WriteHeader(r.StatusCode)
	w.WriteString(string(res))
}