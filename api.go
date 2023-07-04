package main

import (
	"bytes"
	"context"
	"image/png"
	"net/http"
	"strconv"

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

func generateButtonConfig(c *gin.Context) (*ButtonConfig, error) {
	paddingX, err := strconv.ParseInt(c.Query("px"), 10, 32)
	if err != nil {
		paddingX = 8
	}

	paddingY, err := strconv.ParseInt(c.Query("py"), 10, 32)
	if err != nil {
		paddingY = 8
	}

	bgColor := c.DefaultQuery("bg", "#4D4D4D")

	text := c.Query("text")

	fontSize, err := strconv.ParseFloat(c.Query("py"), 64)
	if err != nil {
		fontSize = 16
	}

	fgColor := c.DefaultQuery("fg", "#FFFFFF")

	config := &ButtonConfig{
		PaddingX:        int(paddingX),
		PaddingY:        int(paddingY),
		BackgroundColor: bgColor,
		Text:            text,
		FontSize:        fontSize,
		TextColor:       fgColor,
	}

	return config, nil
}

func (s *ApiServer) handleGetButton(c *gin.Context) {
	buttonConfig, err := generateButtonConfig(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Data",
		})
	}

	button, serviceErr := s.service.GetButton(context.Background(), buttonConfig)

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
