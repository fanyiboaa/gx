package main

import (
	"github.com/fanyiboaa/gx"
	"github.com/gin-gonic/gin"
)

type Request struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func weHandler() (any, error) {
	return gin.H{"message": "hello we"}, nil
}

func wceHandler(ctx *gin.Context) (any, error) {
	return gin.H{"message": "hello wce"}, nil
}

func wHandler(req *Request) (any, error) {
	return gin.H{"message": "hello w", "data": req}, nil
}

func wcHandler(ctx *gin.Context, req *Request) (any, error) {
	return gin.H{"message": "hello wc", "data": req}, nil
}

func main() {
	engine := gin.Default()

	engine.GET("/we", gx.WE(weHandler))
	engine.GET("/wce", gx.WCE(wceHandler))
	engine.POST("/w", gx.W(wHandler))
	engine.POST("/wc", gx.WC(wcHandler))

	engine.Run(":8080")
}
