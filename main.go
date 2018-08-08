package main

import (
	"github.com/piyush2206/go-domain-driven-design/admin"
	"github.com/piyush2206/go-domain-driven-design/dep"
	"github.com/piyush2206/go-domain-driven-design/report"
	"go.uber.org/fx"
)

func main() {
	// fx.New creates a new fx app that can be handled with fx.lifecycle methods
	app := fx.New(
		// external dependencies constructor functions
		dep.FxModule,

		// group of all admin contructor and register functions
		admin.FxModule,

		// group of all report contructor and register functions
		report.FxModule,

		// group of all report contructor and register functions
		gRPCFxModule,
	)

	// Run the app continuously untill an error is recived from any 'invoke' functions
	app.Run()
}
