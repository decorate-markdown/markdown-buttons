package main

import (
	"image"
)

type ButtonRequest struct {
	Text string `json:"text"`
}

type ButtonResponse struct {
	Button image.Image `json:"button"`
}
