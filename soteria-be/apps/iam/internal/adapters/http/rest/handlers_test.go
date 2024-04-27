package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	coretest "github.com/cmclaughlin24/soteria-be/apps/iam/internal/core-test"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

func TestHandler_findPermissions(t *testing.T) {
	tests := []struct {
		name       string
		drivers    *ports.Drivers
		statusCode int
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusOK,
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Drivers{PermissionsService: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.drivers}
			req, _ := http.NewRequest("GET", "/permissions", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.findPermissions)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_findPermission(t *testing.T) {
	tests := []struct {
		name       string
		drivers    *ports.Drivers
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Drivers{PermissionsService: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.drivers}
			req, _ := http.NewRequest("GET", "/permission/"+tt.id, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.findPermission)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}
