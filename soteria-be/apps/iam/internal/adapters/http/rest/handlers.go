package rest

import (
	"net/http"

	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/core/ports"
	"github.com/cmclaughlin24/soteria-be/apps/iam/internal/pkg/auth"
	"github.com/go-chi/chi/v5"
)

type result struct {
	data any
	err  error
}

type Handler struct {
	drivers *ports.Drivers
}

func NewHandler(drivers *ports.Drivers) *Handler {
	return &Handler{
		drivers: drivers,
	}
}

func (h *Handler) findPermissions(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan result)

	go func() {
		permissions, err := h.drivers.PermissionsService.FindAll(r.Context())
		resultChan <- result{permissions, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) findPermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resultChan := make(chan result)

	go func() {
		p, err := h.drivers.PermissionsService.FindOne(r.Context(), id)
		resultChan <- result{p, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) createPermission(w http.ResponseWriter, r *http.Request) {
	var dto CreatePermissionDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		p, err := h.drivers.PermissionsService.Create(r.Context(), *domain.NewPermission(
			"",
			dto.Resource,
			dto.Action,
		))
		resultChan <- result{p, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusCreated, ApiResponseDto{
			Message: "Sucessfully created permission!",
			Data:    res.data,
		})
	}
}

func (h *Handler) updatePermssion(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var dto UpdatePermissionDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		p, err := h.drivers.PermissionsService.Update(r.Context(), *domain.NewPermission(
			id,
			dto.Resource,
			dto.Action,
		))
		resultChan <- result{p, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Succesfully updated permission!",
			Data:    res.data,
		})
	}
}

func (h *Handler) removePermission(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	resultChan := make(chan result)

	go func() {
		err := h.drivers.PermissionsService.Remove(r.Context(), id)
		resultChan <- result{nil, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully deleted permission!",
		})
	}
}

func (h *Handler) findUsers(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan result)

	go func() {
		users, err := h.drivers.UserService.FindAll(r.Context())
		resultChan <- result{users, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) findUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resultChan := make(chan result)

	go func() {
		u, err := h.drivers.UserService.FindOne(r.Context(), id)
		resultChan <- result{u, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var dto CreateUserDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	u := domain.NewUser(
		"",
		dto.Name,
		dto.Email,
		dto.PhoneNumber,
		dto.Password,
		dto.DeliveryMethods,
		dto.TimeZone,
	)

	for _, permissionDto := range dto.Permissions {
		u.AddPermission(domain.UserPermission{
			Resource: permissionDto.Resource,
			Action:   permissionDto.Action,
		})
	}

	resultChan := make(chan result)

	go func() {
		u, err := h.drivers.UserService.Create(r.Context(), *u)
		resultChan <- result{u, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusCreated, ApiResponseDto{
			Message: "Sucessfully created user!",
			Data:    res.data,
		})
	}
}

func (h *Handler) updateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var dto UpdateUserDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	u := domain.NewUser(
		id,
		dto.Name,
		dto.Email,
		dto.PhoneNumber,
		"",
		dto.DeliveryMethods,
		dto.TimeZone,
	)

	for _, permissionDto := range dto.Permissions {
		u.AddPermission(domain.UserPermission{
			Resource: permissionDto.Resource,
			Action:   permissionDto.Action,
		})
	}

	resultChan := make(chan result)

	go func() {
		u, err := h.drivers.UserService.Update(r.Context(), *u)
		resultChan <- result{u, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusCreated, ApiResponseDto{
			Message: "Sucessfully updated user!",
			Data:    res.data,
		})
	}
}

func (h *Handler) removeUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resultChan := make(chan result)

	go func() {
		err := h.drivers.UserService.Remove(r.Context(), id)
		resultChan <- result{nil, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully deleted user!",
		})
	}
}

func (h *Handler) createApiKey(w http.ResponseWriter, r *http.Request) {
	var dto CreateApiKeyDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	permissions := make([]domain.UserPermission, 0, len(dto.Permissions))

	for _, permissionDto := range dto.Permissions {
		permissions = append(permissions, domain.UserPermission{
			Resource: permissionDto.Resource,
			Action:   permissionDto.Action,
		})
	}

	sub, _ := auth.ClaimsFromContext(r.Context()).GetSubject()
	resultChan := make(chan result)

	go func() {
		apiKey, err := h.drivers.ApiKeyService.Create(r.Context(), dto.Name, permissions, sub)
		resultChan <- result{apiKey, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusCreated, ApiResponseDto{
			Message: "Sucessfully created api key!",
			Data:    res.data,
		})
	}
}

func (h *Handler) removeApiKey(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	resultChan := make(chan result)

	go func() {
		err := h.drivers.ApiKeyService.Remove(r.Context(), id)
		resultChan <- result{nil, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully deleted user!",
		})
	}
}

func (h *Handler) verifyApiKey(w http.ResponseWriter, r *http.Request) {
	var dto VerifyApiKeyDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		claims, err := h.drivers.ApiKeyService.VerifyApiKey(r.Context(), dto.Key)
		resultChan <- result{claims, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully authenicated!",
			Data:    res.data,
		})
	}
}

func (h *Handler) signup(w http.ResponseWriter, r *http.Request) {
	var dto SignUpDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		_, err := h.drivers.UserService.Create(r.Context(), *domain.NewUser(
			"",
			dto.Name,
			dto.Email,
			dto.PhoneNumber,
			dto.Password,
			[]string{},
			"",
		))
		resultChan <- result{nil, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusCreated, ApiResponseDto{
			Message: "Successfully signed up!",
		})
	}
}

func (h *Handler) signin(w http.ResponseWriter, r *http.Request) {
	var dto SignInDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		tokens, err := h.drivers.AuthenticationService.Signin(r.Context(), dto.Email, dto.Password)
		resultChan <- result{tokens, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusInternalServerError, ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully logged in!",
			Data:    res.data,
		})
	}
}

func (h *Handler) verifyAccessToken(w http.ResponseWriter, r *http.Request) {
	var dto TokenDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		claims, err := h.drivers.AuthenticationService.VerifyAccessToken(r.Context(), dto.Token)
		resultChan <- result{claims, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusUnauthorized, ApiErrorResponseDto{
				Message:    "Invalid access token",
				Error:      "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully authenicated!",
			Data:    res.data,
		})
	}
}

func (h *Handler) refreshAccessToken(w http.ResponseWriter, r *http.Request) {
	var dto TokenDto

	if err := readJsonPayload(r, &dto); err != nil {
		sendJsonResponse(w, http.StatusBadRequest, ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	resultChan := make(chan result)

	go func() {
		tokens, err := h.drivers.AuthenticationService.RefreshAccessToken(r.Context(), dto.Token)
		resultChan <- result{tokens, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			sendJsonResponse(w, http.StatusUnauthorized, ApiErrorResponseDto{
				Message:    "Invalid refresh token",
				Error:      "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			})
			return
		}

		sendJsonResponse(w, http.StatusOK, ApiResponseDto{
			Message: "Successfully refreshed tokens!",
			Data:    res.data,
		})
	}
}
