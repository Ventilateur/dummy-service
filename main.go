package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Response struct {
	RequestHeader http.Header `json:"requestHeader"`
	RemoteAddr    string      `json:"remoteAddr"`
}

type ErrorResponse = struct {
	Message string `json:"message"`
}

func handler(c *gin.Context) {
	code, err := strconv.Atoi(c.Param("code"))
	if err != nil || code < 100 || code > 599 {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			ErrorResponse{Message: "Request path must start with a valid HTTP code"},
		)
	}

	c.JSON(code, &Response{
		RequestHeader: c.Request.Header,
		RemoteAddr:    c.ClientIP(),
	})
}

func main() {
	r := gin.Default()
	r.GET("/:code/*path", handler)
	r.HEAD("/:code/*path", handler)
	r.POST("/:code/*path", handler)
	r.PUT("/:code/*path", handler)
	r.DELETE("/:code/*path", handler)
	r.OPTIONS("/:code/*path", handler)
	r.PATCH("/:code/*path", handler)
	log.Fatal(r.Run())
}
