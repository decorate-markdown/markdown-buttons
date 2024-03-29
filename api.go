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

const (
	defaultBGColor  string = "#4D4D4D"
	defaultFGColor  string = "#FFFFFF"
	defaultFont     string = "RobotoMono-Regular"
	defaultPaddingX int    = 12
	defaultPaddingY int    = 12
	defaultFontSize int    = 14
	defaultText     string = ""
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{service}
}

func (s *ApiServer) Start() {
	r := gin.Default()
	r.GET("/btn", s.handleGetButton)
	r.Run()
}

func generateButtonConfig(c *gin.Context) (*ButtonConfig, error) {
	paddingX, err := strconv.ParseInt(c.Query("px"), 10, 32)
	if err != nil {
		paddingX = int64(defaultPaddingX)
	}

	paddingY, err := strconv.ParseInt(c.Query("py"), 10, 32)
	if err != nil {
		paddingY = int64(defaultPaddingY)
	}

	bgColor := c.DefaultQuery("bg", defaultBGColor)

	fontName := c.DefaultQuery("font", defaultFont)

	text := c.DefaultQuery("text", defaultText)

	fontSize, err := strconv.ParseInt(c.Query("size"), 10, 32)
	if err != nil {
		fontSize = int64(defaultFontSize)
	}

	fgColor := c.DefaultQuery("fg", defaultFGColor)

	config := &ButtonConfig{
		PaddingX:        int(paddingX),
		PaddingY:        int(paddingY),
		BackgroundColor: bgColor,
		FontName:        fontName,
		Text:            text,
		FontSize:        int(fontSize),
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
