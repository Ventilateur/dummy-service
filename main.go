package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	path = "/rest"
	msg = "default"
)

type Response struct {
	RequestHeader http.Header`json:"request_header"`
	Message string
}
func init() {
	if p := os.Getenv("LISTEN_PATH"); p != "" {
		path = p
	}
	if m := os.Getenv("RESP_MSG"); m != "" {
		msg = m
	}
}

func main() {
	r := gin.Default()
	r.GET(path, func(c *gin.Context) {
		c.JSON(200, &Response{
			RequestHeader: c.Request.Header,
			Message: msg,
		})
	})
	log.Fatal(r.Run())
}
