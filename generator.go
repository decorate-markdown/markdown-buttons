package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 9:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x%02x", &c.R, &c.G, &c.B, &c.A)
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	default:
		err = fmt.Errorf("invalid length, must be 9 or 7")
	}
	return
}

func addLabel(img *image.RGBA, x, y int, label string, col color.RGBA, fontScale float32) {
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}

	d.DrawString(label)
}

func GenerateButton(config *ButtonConfig) (image.Image, error) {
	fontWidth := 7 * config.FontScale
	fontHeight := 9 * config.FontScale

	width := int(fontWidth)*len(config.Text) + (2 * config.PaddingX)
	height := int(fontHeight) + (2 * config.PaddingY)

	button := image.NewRGBA(image.Rect(
		0,
		0,
		width,
		height,
	))

	bgCol, err := ParseHexColor(config.BackgroundColor)
	if err != nil {
		return nil, err
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			button.Set(x, y, bgCol)
		}
	}

	textCol, err := ParseHexColor(config.TextColor)
	if err != nil {
		return nil, err
	}

	addLabel(button, config.PaddingX, 9+config.PaddingY, config.Text, textCol, config.FontScale)

	return button, nil
}
