package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type ButtonConfig struct {
	CONTENT []string `json:"content"`
}

func main() {
	fmt.Print("Hello World!")

	r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "WORKS",
    })
  })

  r.Run()
}