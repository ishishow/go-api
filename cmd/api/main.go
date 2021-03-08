package main

import (
	"20dojo-online/pkg/infrastructure/waf"
)

func main() {
	// if you want to use grpc-gateway use 8081 port and $ make gateway
	// serverPort := "8081"
	serverPort := "50051"
	waf.Serve(serverPort)
}
