package main

import (
	"bytes"
	"context"
	"image/png"
	"net/http"
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{service}
}

func (s *ApiServer) Start(listenAddr string) error {
	http.HandleFunc("/", s.handleGetButton)
	return http.ListenAndServe(listenAddr, nil)
}

func (s *ApiServer) handleGetButton(w http.ResponseWriter, r *http.Request) {
	button, serviceErr := s.service.GetButton(context.Background(), &ButtonRequest{Text: "Hello"})

	if serviceErr != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(serviceErr.Error()))
	}

	buf := new(bytes.Buffer)
	encodeErr := png.Encode(buf, button.Button)

	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Add("Content-Type", "text/plain")
		w.Write([]byte(encodeErr.Error()))
	}

	imageBytes := buf.Bytes()

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "image/png")
	w.Write(imageBytes)
}
