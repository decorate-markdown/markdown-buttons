package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	service := NewLoggingService(NewMarkdownButtonsService())

	req := ButtonRequest{Text: "Hello World!"}

	button, err := service.GetButton(context.TODO(), &req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", button)
}
