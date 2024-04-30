package rest

import (
	"net/http"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(drivers *ports.Drivers) http.Handler {
	handler := NewHandler(drivers)
	mux := chi.NewRouter()
	accessTokenVerifier := iam.AuthenticateAccessToken(drivers.AuthenticationService)
	apiKeyVerifier := iam.AuthenticateApiKey(drivers.AuthenticationService)

	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(1000 * time.Millisecond))

	mux.Route("/facilities", func(r chi.Router) {
		r.Use(iam.Authenticate(accessTokenVerifier, apiKeyVerifier))

		r.With(iam.Authorize("facility", "list")).Get("/", handler.findFacilities)
		r.With(iam.Authorize("facility", "get")).Get("/{code}", handler.findFacility)
		r.With(iam.Authorize("facility", "create")).Post("/", handler.createFacility)
		r.With(iam.Authorize("facility", "update")).Patch("/{code}", handler.updateFacility)
		r.With(iam.Authorize("facility", "remove")).Delete("/{code}", handler.removeFacility)
	})

	return mux
}
