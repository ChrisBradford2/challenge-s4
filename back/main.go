package main

import (
	"github.com/gin-gonic/gin"
	"github.com/madkins23/gin-utils/pkg/ginzero"
)

func main() {
	r := gin.New()
	r.Use(ginzero.Logger())

	r.GET("/", func(c *gin.Context) {
		c.String(200, "hello, gin-zerolog example")
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.Run(":8080")
}
