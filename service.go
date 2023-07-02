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
	button := image.NewRGBA(image.Rect(0, 0, 100, 100))

	return &ButtonResponse{Button: button}, nil
}
