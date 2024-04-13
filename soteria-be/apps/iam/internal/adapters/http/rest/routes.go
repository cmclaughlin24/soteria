package rest

import (
	"time"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(drivers *ports.Drivers) *chi.Mux {
	handler := NewHandler(drivers)
	mux := chi.NewRouter()
	accessTokenVerifier := AuthenticateAccessToken(drivers.AuthenticationService)
	apiKeyVerifier := AuthenticateApiKey(drivers.ApiKeyService)

	mux.Use(middleware.Logger)
	mux.Use(middleware.Timeout(1000 * time.Millisecond))

	mux.Route("/permissions", func(r chi.Router) {
		r.Use(Authenticate(accessTokenVerifier, apiKeyVerifier))

		r.With(Authorize("permission", "list")).Get("/", handler.findPermissions)
		r.With(Authorize("permission", "get")).Get("/{id}", handler.findPermission)
		r.With(Authorize("permission", "create")).Post("/", handler.createPermission)
		r.With(Authorize("permission", "update")).Patch("/{id}", handler.updatePermssion)
		r.With(Authorize("permission", "remove")).Delete("/{id}", handler.removePermission)
	})

	mux.Route("/users", func(r chi.Router) {
		r.Use(Authenticate(accessTokenVerifier, apiKeyVerifier))

		r.With(Authorize("user", "list")).Get("/", handler.findUsers)
		r.With(Authorize("user", "get")).Get("/{id}", handler.findUser)
		r.With(Authorize("user", "create")).Post("/", handler.createUser)
		r.With(Authorize("user", "update")).Patch("/{id}", handler.updateUser)
		r.With(Authorize("user", "remove")).Delete("/{id}", handler.removeUser)
	})

	mux.Route("/api-keys", func(r chi.Router) {
		r.With(Authenticate(accessTokenVerifier, apiKeyVerifier), Authorize("api_key", "create")).Post("/", handler.createApiKey)
		r.With(Authenticate(accessTokenVerifier, apiKeyVerifier), Authorize("api_key", "remove")).Delete("/{id}", handler.removeApiKey)
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
