package rest

import (
	"net/http"
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(drivers *ports.Drivers) http.Handler {
	handler := NewHandler(drivers)
	mux := chi.NewRouter()
	accessTokenVerifier := iam.AuthenticateAccessToken(drivers.AuthenticationService)
	apiKeyVerifier := iam.AuthenticateApiKey(drivers.ApiKeyService)

	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(1000 * time.Millisecond))

	mux.Route("/permissions", func(r chi.Router) {
		r.Use(iam.Authenticate(accessTokenVerifier, apiKeyVerifier))

		r.With(iam.Authorize("permission", "list")).Get("/", handler.findPermissions)
		r.With(iam.Authorize("permission", "get")).Get("/{id}", handler.findPermission)
		r.With(iam.Authorize("permission", "create")).Post("/", handler.createPermission)
		r.With(iam.Authorize("permission", "update")).Patch("/{id}", handler.updatePermssion)
		r.With(iam.Authorize("permission", "remove")).Delete("/{id}", handler.removePermission)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Use(iam.Authenticate(accessTokenVerifier, apiKeyVerifier))

		r.With(iam.Authorize("user", "list")).Get("/", handler.findUsers)
		r.With(iam.Authorize("user", "get")).Get("/{id}", handler.findUser)
		r.With(iam.Authorize("user", "create")).Post("/", handler.createUser)
		r.With(iam.Authorize("user", "update")).Patch("/{id}", handler.updateUser)
		r.With(iam.Authorize("user", "remove")).Delete("/{id}", handler.removeUser)
	})

	mux.Route("/api-keys", func(r chi.Router) {
		r.With(iam.Authenticate(accessTokenVerifier, apiKeyVerifier), iam.Authorize("api_key", "create")).Post("/", handler.createApiKey)
		r.With(iam.Authenticate(accessTokenVerifier, apiKeyVerifier), iam.Authorize("api_key", "remove")).Delete("/{id}", handler.removeApiKey)
		r.Post("/verify", handler.verifyApiKey)
	})

	mux.Route("/authentication", func(r chi.Router) {
		r.Post("/sign-up", handler.signup)
		r.Post("/sign-in", handler.signin)
		r.Post("/verify", handler.verifyAccessToken)
		r.Post("/refresh", handler.refreshAccessToken)
	})

	return mux
}
