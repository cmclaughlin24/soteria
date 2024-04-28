package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/adapters/http/rest"
)

func main() {
	httpPort := os.Getenv("HTTP_PORT")

	mux := rest.Routes()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", httpPort), mux); err != nil {
		panic(err)
	}
}
