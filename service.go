package main

import (
	"context"
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
