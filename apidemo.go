package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/info", Info)
	r.Run(":10088")
}

func Info(c *gin.Context) {
	currentTime := time.Now().Format(time.RFC1123)
	c.JSON(http.StatusOK, gin.H{
		"version": "1.5.3",
		"project": "hello world",
		"name": "YH",
		"timestamp":  currentTime,
	})
}
