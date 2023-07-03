package main

func main() {
	service := NewLoggingService(NewMarkdownButtonsService())

	apiServer := NewApiServer(service)
	apiServer.Start()
}
