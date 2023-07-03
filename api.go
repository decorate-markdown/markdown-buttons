package main

import (
	"bytes"
	"context"
	"image/png"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{service}
}

func (s *ApiServer) Start() {
	r := gin.Default()

	r.GET("/", s.handleGetButton)

	r.Run()
}

func (s *ApiServer) handleGetButton(c *gin.Context) {
	button, serviceErr := s.service.GetButton(context.Background(), &ButtonRequest{Text: "Hello"})

	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Service Error",
		})
	}

	buf := new(bytes.Buffer)
	encodeErr := png.Encode(buf, button.Button)

	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Data",
		})
	}

	imageBytes := buf.Bytes()

	c.Render(http.StatusOK, render.Data{Data: imageBytes})
}
