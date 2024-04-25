package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/grpc"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/http/rest"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	grpcPort := os.Getenv("GRPC_PORT")
	drivers, err := core.Init()

	if err != nil {
		panic(err)
	}

	mux := rest.Routes(drivers)
	grpcServer := grpc.NewIamGRPCServer(drivers)
	go grpcServer.Listen(grpcPort)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux); err != nil {
		panic(err)
	}
}
