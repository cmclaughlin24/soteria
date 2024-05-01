package rest

import (
	"net/http"
	"strings"
	"testing"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/go-chi/chi/v5"
)

func TestRoutes(t *testing.T) {
	// Arrange.
	routes := []struct {
		route  string
		method string
	}{
		{"/api-keys/", "POST"},
		{"/api-keys/{id}", "DELETE"},
		{"/api-keys/verify", "POST"},
		{"/authentication/sign-up", "POST"},
		{"/authentication/sign-in", "POST"},
		{"/authentication/verify", "POST"},
		{"/authentication/refresh", "POST"},
		{"/permissions/", "GET"},
		{"/permissions/{id}", "GET"},
		{"/permissions/", "POST"},
		{"/permissions/{id}", "PATCH"},
		{"/permissions/{id}", "DELETE"},
		{"/users/", "GET"},
		{"/users/{id}", "GET"},
		{"/users/", "POST"},
		{"/users/{id}", "PATCH"},
		{"/users/{id}", "DELETE"},
	}
	mux := Routes(&ports.Services{})

	for _, r := range routes {
		// Act/Assert.
		if !routeExists(r.route, r.method, mux.(chi.Routes)) {
			t.Errorf("expected route [%s] %s to exist", r.method, r.route)
		}
	}
}

func routeExists(route, method string, routes chi.Routes) bool {
	var found bool

	chi.Walk(routes, func(m string, r string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		if strings.EqualFold(r, route) && strings.EqualFold(m, method) {
			found = true
		}

		return nil
	})

	return found
}
