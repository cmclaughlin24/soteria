package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	coretest "github.com/cmclaughlin24/soteria-be/apps/iam/internal/core-test"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
)

func TestHandler_findPermissions(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusOK,
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Permission: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
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
		services   *ports.Services
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Permission: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
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
		services   *ports.Services
		statusCode int
		dto        CreatePermissionDto
	}{
		{
			"should yield an CREATED status code if the request was successful",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusCreated,
			CreatePermissionDto{"australian-cattle-dog", "adopt"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusBadRequest,
			CreatePermissionDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Permission: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			CreatePermissionDto{"australian-cattle-dog", "adopt"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
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
		services   *ports.Services
		statusCode int
		id         string
		dto        UpdatePermissionDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusOK,
			"1",
			UpdatePermissionDto{"australian-cattle-dog", "train"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusBadRequest,
			"1",
			UpdatePermissionDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Permission: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			"1",
			UpdatePermissionDto{"australian-cattle-dog", "train"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
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

func TestHandler_removePermission(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Permission: coretest.NewSuccessPermissionService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Permission: coretest.NewErrorPermissionService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			req, _ := http.NewRequest("DELETE", "/permission/"+tt.id, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.removePermission)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_findUsers(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusOK,
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			req, _ := http.NewRequest("GET", "/users", nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.findUsers)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_findUser(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			req, _ := http.NewRequest("GET", "/users/"+tt.id, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.findUser)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_createUser(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        CreateUserDto
	}{
		{
			"should yield an CREATED status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusCreated,
			CreateUserDto{
				"sydney",
				"sydney@aaustralian-cattle-dog.com",
				"+2815362118",
				"hduson",
				"",
				nil,
				nil,
			},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusBadRequest,
			CreateUserDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
			CreateUserDto{
				"sydney",
				"sydney@aaustralian-cattle-dog.com",
				"+2815362118",
				"hduson",
				"",
				nil,
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/users", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.createUser)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_updateUser(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		id         string
		dto        UpdateUserDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusOK,
			"1",
			UpdateUserDto{
				"hudson",
				"hudson@aaustralian-cattle-dog.com",
				"+2815362118",
				"",
				nil,
				nil,
			},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusBadRequest,
			"1",
			UpdateUserDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
			"1",
			UpdateUserDto{
				"hudson",
				"hudson@aaustralian-cattle-dog.com",
				"+2815362118",
				"",
				nil,
				nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("PATCH", "/users/"+tt.id, bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.updateUser)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_removeUser(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			req, _ := http.NewRequest("DELETE", "/users/"+tt.id, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.removeUser)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_createApiKey(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        CreateApiKeyDto
	}{
		{
			"should yield an CREATED status code if the request was successful",
			&ports.Services{ApiKey: coretest.NewSuccessApiKeyService()},
			http.StatusCreated,
			CreateApiKeyDto{"australian-cattle-dog", []UserPermissionDto{{"training", "sit"}}},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{ApiKey: coretest.NewSuccessApiKeyService()},
			http.StatusBadRequest,
			CreateApiKeyDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{ApiKey: coretest.NewErrorApiKeyService()},
			http.StatusInternalServerError,
			CreateApiKeyDto{"australian-cattle-dog", []UserPermissionDto{{"training", "sit"}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			claims := iam.AccessTokenClaims{}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequestWithContext(iam.SetContext(context.Background(), claims), "POST", "/api-keys", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.createApiKey)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_removeApiKey(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		id         string
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{ApiKey: coretest.NewSuccessApiKeyService()},
			http.StatusOK,
			"successful",
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{ApiKey: coretest.NewErrorApiKeyService()},
			http.StatusInternalServerError,
			"error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			claims := iam.AccessTokenClaims{}
			req, _ := http.NewRequestWithContext(iam.SetContext(context.Background(), claims), "DELETE", "/api-keys/"+tt.id, nil)
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.removeApiKey)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_verifyApiKey(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        VerifyApiKeyDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{ApiKey: coretest.NewSuccessApiKeyService()},
			http.StatusOK,
			VerifyApiKeyDto{"the-oldest-cattled-dog-lived-to-twenty-nine"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{ApiKey: coretest.NewSuccessApiKeyService()},
			http.StatusBadRequest,
			VerifyApiKeyDto{},
		},
		{
			"should yield an UNAUTHORIZED status code if the request fails",
			&ports.Services{ApiKey: coretest.NewErrorApiKeyService()},
			http.StatusUnauthorized,
			VerifyApiKeyDto{"the-oldest-cattled-dog-lived-to-twenty-nine"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			claims := iam.AccessTokenClaims{}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequestWithContext(iam.SetContext(context.Background(), claims), "POST", "/api-keys", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.verifyApiKey)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_signup(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        SignUpDto
	}{
		{
			"should yield an CREATED status code if the request was successful",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusCreated,
			SignUpDto{
				"sydney",
				"sydney@aaustralian-cattle-dog.com",
				"+2815362118",
				"hduson",
			},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{User: coretest.NewSuccessUserService()},
			http.StatusBadRequest,
			SignUpDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{User: coretest.NewErrorUserService()},
			http.StatusInternalServerError,
			SignUpDto{
				"sydney",
				"sydney@aaustralian-cattle-dog.com",
				"+2815362118",
				"hduson",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/sign-up", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.signup)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_signin(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        SignInDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusOK,
			SignInDto{"sydney@aaustralian-cattle-dog.com", "hduson"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusBadRequest,
			SignInDto{},
		},
		{
			"should yield an INTERNAL SERVER ERROR status code if the request fails",
			&ports.Services{Authentication: coretest.NewErrorAuthenticationService()},
			http.StatusInternalServerError,
			SignInDto{"sydney@aaustralian-cattle-dog.com", "hduson"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/sign-in", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.signin)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_verifyAccessToken(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        TokenDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusOK,
			TokenDto{"cattle-dogs-are-hard-working-and-loyal"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusBadRequest,
			TokenDto{},
		},
		{
			"should yield an UNAUTHORIZED status code if the request fails",
			&ports.Services{Authentication: coretest.NewErrorAuthenticationService()},
			http.StatusUnauthorized,
			TokenDto{"cattle-dogs-are-hard-working-and-loyal"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/verify", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.verifyAccessToken)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestHandler_refreshAccessToken(t *testing.T) {
	tests := []struct {
		name       string
		services   *ports.Services
		statusCode int
		dto        TokenDto
	}{
		{
			"should yield an OK status code if the request was successful",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusOK,
			TokenDto{"cattle-dogs-are-hard-working-and-loyal"},
		},
		{
			"should yield a BAD REQUEST status code if the payload is invalid",
			&ports.Services{Authentication: coretest.NewSuccessAuthenticationService()},
			http.StatusBadRequest,
			TokenDto{},
		},
		{
			"should yield an UNAUTHORIZED status code if the request fails",
			&ports.Services{Authentication: coretest.NewErrorAuthenticationService()},
			http.StatusUnauthorized,
			TokenDto{"cattle-dogs-are-hard-working-and-loyal"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange.
			h := &Handler{tt.services}
			body, _ := json.Marshal(tt.dto)
			req, _ := http.NewRequest("POST", "/refresh", bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(h.refreshAccessToken)

			// Act.
			handler.ServeHTTP(rr, req)

			// Assert.
			if rr.Code != tt.statusCode {
				t.Errorf("expected status code %d but received %d", tt.statusCode, rr.Code)
			}
		})
	}
}
