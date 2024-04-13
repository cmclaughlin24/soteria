package main

import (
	"net/http"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/adapters/http/rest"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core"
)

func main() {
	drivers, err := core.Init()

	if err != nil {
		panic(err)
	}

	mux := rest.Routes(drivers)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
