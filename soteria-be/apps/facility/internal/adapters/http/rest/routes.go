package rest

import (
	"net/http"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(drivers *ports.Drivers) http.Handler {
	handler := NewHandler(drivers)
	mux := chi.NewRouter()

	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(1000 * time.Millisecond))

	mux.Route("/facilities", func(r chi.Router) {
		r.Get("/", handler.findFacilities)
		r.Get("/{code}", handler.findFacility)
		r.Post("/", handler.createFacility)
		r.Patch("/{code}", handler.updateFacility)
		r.Delete("/{code}", handler.removeFacility)
	})

	return mux
}
