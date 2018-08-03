package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/piyush2206/go-domain-driven-design/admin"
	"github.com/piyush2206/go-domain-driven-design/report"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = "9090"
	httpPort = "8080"
)

func startGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// registering admin grpc servers
	admin.RegisterClassServer(s, new(admin.ClassService))
	admin.RegisterStudentServer(s, new(admin.StudentService))

	// registering report grpc servers
	report.RegisterReportServer(s, new(report.ReportService))

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGRPCGateway() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	grpcAddr := fmt.Sprintf("localhost:%s", grpcPort)

	for _, f := range []func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error{
		admin.RegisterClassHandlerFromEndpoint,
		admin.RegisterStudentHandlerFromEndpoint,

		report.RegisterReportHandlerFromEndpoint,
	} {
		if err := f(ctx, mux, grpcAddr, opts); err != nil {
			return err
		}
	}

	return http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux)
}
