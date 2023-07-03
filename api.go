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
	button, serviceErr := s.service.GetButton(context.Background(), &ButtonConfig{
		PaddingX:        4,
		PaddingY:        4,
		BackgroundColor: "#282828",
		Text:            "CHECK CHECK CHECK CHECK",
		FontScale:       1.1,
		TextColor:       "#ebdbb2",
	})

	if serviceErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Service Error",
		})
	}

	buf := new(bytes.Buffer)
	encodeErr := png.Encode(buf, *button)

	if encodeErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Data",
		})
	}

	imageBytes := buf.Bytes()

	c.Writer.Header().Set("Content-Type", "image/png")

	c.Render(http.StatusOK, render.Data{Data: imageBytes})
}
