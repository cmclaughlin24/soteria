package rest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
)

func TestHandler_findPermissions(t *testing.T) {
	// Fixme: Decide how to implement coretest package for unit testing.
	tests := []struct {
		name       string
		drivers    *ports.Drivers
		statusCode int
	}{
		{"first test", &ports.Drivers{}, http.StatusOK},
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
