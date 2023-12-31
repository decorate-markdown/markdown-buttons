package main

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
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

func setBackground(img *image.RGBA, width, height int, col string) error {
	bgCol, err := ParseHexColor(col)
	if err != nil {
		return err
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, bgCol)
		}
	}

	return nil
}

func initializeFont(fontName string) (*truetype.Font, error) {
	fontBytes, err := ioutil.ReadFile("fonts/" + fontName + ".ttf")
	if err != nil {
		return nil, err
	}

	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return nil, err
	}

	return font, nil
}

func setupFont(img *image.RGBA, fontName string, fontSize int, col string) (*freetype.Context, error) {
	font, err := initializeFont(fontName)
	if err != nil {
		return nil, err
	}

	textCol, err := ParseHexColor(col)
	if err != nil {
		return nil, err
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(float64(fontSize))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(textCol))

	return c, nil
}

func addLabel(img *image.RGBA, x, y int, label string, col string, fontName string, fontSize int) error {
	c, err := setupFont(img, fontName, fontSize, col)
	if err != nil {
		return err
	}

	pt := freetype.Pt(x, y)

	c.DrawString(label, pt)

	return nil
}

func GenerateButton(config *ButtonConfig) (image.Image, error) {
	f, err := initializeFont(config.FontName)
	if err != nil {
		return nil, err
	}

	opts := truetype.Options{}
	opts.Size = float64(config.FontSize)
	face := truetype.NewFace(f, &opts)

	textWidth := font.MeasureString(face, config.Text).Ceil()
	textHeight := face.Metrics().Height.Ceil() - int(config.FontSize/4)

	width := textWidth + (2 * config.PaddingX)
	height := textHeight + (2 * config.PaddingY)

	button := image.NewRGBA(image.Rect(0, 0, width, height))

	setBackground(button, width, height, config.BackgroundColor)
	addLabel(button, config.PaddingX, textHeight+config.PaddingY, config.Text, config.TextColor, config.FontName, config.FontSize)

	return button, nil
}
