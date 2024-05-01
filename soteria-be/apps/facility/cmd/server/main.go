package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/adapters/http/rest"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")
	services, err := core.Init()

	if err != nil {
		panic(err)
	}

	mux := rest.Routes(services)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux); err != nil {
		panic(err)
	}
}
