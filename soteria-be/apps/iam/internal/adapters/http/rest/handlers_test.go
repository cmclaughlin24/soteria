package rest

import (
	"bytes"
	"encoding/json"
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

func TestHandler_createPermission(t *testing.T) {
	tests := []struct {
		name       string
		drivers    *ports.Drivers
		statusCode int
		dto        CreatePermissionDto
	}{
		{
			"should yield an CREATED status code if the request was successful",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusCreated,
			CreatePermissionDto{"australian-cattle-dog", "adopt"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusBadRequest,
			CreatePermissionDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Drivers{PermissionsService: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			CreatePermissionDto{"australian-cattle-dog", "adopt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.drivers}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/permission", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.createPermission)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_updatePermssion(t *testing.T) {
	tests := []struct {
		name       string
		drivers    *ports.Drivers
		statusCode int
		id         string
		dto        UpdatePermissionDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusOK,
			"1",
			UpdatePermissionDto{"australian-cattle-dog", "train"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Drivers{PermissionsService: coretest.NewSuccessPermissionService()},
			http.StatusBadRequest,
			"1",
			UpdatePermissionDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Drivers{PermissionsService: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			"1",
			UpdatePermissionDto{"australian-cattle-dog", "train"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.drivers}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("PATCH", "/permission/"+tt.id, bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.updatePermssion)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}
