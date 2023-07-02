package main

import "log"

func main() {
	service := NewLoggingService(NewMarkdownButtonsService())

	apiServer := NewApiServer(service)
	log.Fatal(apiServer.Start(":3000"))
}
