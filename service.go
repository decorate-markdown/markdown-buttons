package main

import (
	"context"
	"image"
)

type Service interface {
	GetButton(context.Context, *ButtonRequest) (*ButtonResponse, error)
}

type MarkdownButtonsService struct{}

func NewMarkdownButtonsService() *MarkdownButtonsService {
	return &MarkdownButtonsService{}
}

func (s *MarkdownButtonsService) GetButton(ctx context.Context, req *ButtonRequest) (*ButtonResponse, error) {
	button, err := GenerateButton(req)

	return &ButtonResponse{Button: button}, err
}

func GenerateButton(config *ButtonRequest) (image.Image, error) {
	button := image.NewRGBA(image.Rect(0, 0, 200, 30))

	return button, nil
}
