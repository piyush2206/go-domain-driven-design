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
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = "9090"
	httpPort = "8080"
)

var gRPCFxModule = fx.Options(
	fx.Provide(
		grpc.NewServer,
	),
	fx.Invoke(
		runGRPC,
	),
)

func runGRPC(s *grpc.Server) {
	go startGRPCGateway()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startGRPCGateway() error {
	ctx := context.Background()
	// ctx, cancel := context.WithCancel(ctx)
	// defer cancel()

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
