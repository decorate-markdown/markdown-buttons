package main

import (
	"context"
	"image"
)

type Service interface {
	GetButton(context.Context, *ButtonRequest) (*image.Image, error)
}

type MarkdownButtonsService struct{}

func NewMarkdownButtonsService() *MarkdownButtonsService {
	return &MarkdownButtonsService{}
}

func (s *MarkdownButtonsService) GetButton(ctx context.Context, req *ButtonRequest) (*image.Image, error) {
	button, err := GenerateButton(req)

	return &button, err
}
