package main

import (
	"github.com/piyush2206/go-domain-driven-design/admin"
	"github.com/piyush2206/go-domain-driven-design/dep"
	"github.com/piyush2206/go-domain-driven-design/report"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

func main() {
	go startGRPCGateway()

	app := fx.New(
		fx.Provide(
			// external dependencies initialisation functions
			dep.NewMySQLConn,

			// GRPC server initialisation
			grpc.NewServer,

			// Class domain initialisation
			admin.NewClassRepository,
			admin.NewClassService,

			// Student domain initialisation
			admin.NewStudentRepository,
			admin.NewStudentService,

			// Report domain initialisation
			report.NewReportService,
		),

		fx.Invoke(
			// Register GRPC services
			admin.RegisterClassServer,
			admin.RegisterStudentServer,
			report.RegisterReportServer,

			// Start GRPC Server
			runGRPC,
		),
	)

	app.Run()
}
