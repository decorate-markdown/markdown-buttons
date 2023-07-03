package main

import "image"

func GenerateButton(config *ButtonRequest) (image.Image, error) {
	button := image.NewRGBA(image.Rect(0, 0, 200, 30))

	return button, nil
}
