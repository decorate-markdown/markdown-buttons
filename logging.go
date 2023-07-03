package main

import (
	"context"
	"fmt"
	"image"
	"time"
)

type LoggingService struct {
	service Service
}

func NewLoggingService(service Service) *LoggingService {
	return &LoggingService{service}
}

func (s *LoggingService) GetButton(ctx context.Context, req *ButtonRequest) (button *image.Image, err error) {
	defer func(start time.Time) {
		fmt.Printf("[%v] %v | Error: %v (Took %v)\n", start, req, err, time.Since(start))
	}(time.Now())

	return s.service.GetButton(ctx, req)
}
