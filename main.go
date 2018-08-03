package main

import (
	"github.com/piyush2206/go-domain-driven-design/admin"
	"github.com/piyush2206/go-domain-driven-design/app"
)

func main() {
	forever := make(chan struct{})

	// Initialize connections to multiple external dependency of the project
	appCtx, err := app.Init()
	if err != nil {
		panic(err)
	}

	admin.Init(appCtx)

	go startGRPCServer()
	go startGRPCGateway()

	<-forever
}
