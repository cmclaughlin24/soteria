package rest

import (
	"net/http"

	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/domain"
	"github.com/cmclaughlin24/soteria-be/apps/facility/internal/core/ports"
	httputils "github.com/cmclaughlin24/soteria-be/pkg/http-utils"
	"github.com/cmclaughlin24/soteria-be/pkg/iam"
	"github.com/go-chi/chi/v5"
)

type result struct {
	data any
	err  error
}

type Handler struct {
	services *ports.Services
}

func NewHandler(services *ports.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) findFacilities(w http.ResponseWriter, r *http.Request) {
	resultChan := make(chan result)

	go func() {
		facilities, err := h.services.Facility.FindAll(r.Context())
		resultChan <- result{facilities, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			httputils.SendJsonResponse(w, http.StatusInternalServerError, httputils.ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		httputils.SendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) findFacility(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	resultChan := make(chan result)

	go func() {
		facility, err := h.services.Facility.FindOne(r.Context(), code)
		resultChan <- result{facility, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			httputils.SendJsonResponse(w, http.StatusInternalServerError, httputils.ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		httputils.SendJsonResponse(w, http.StatusOK, res.data)
	}
}

func (h *Handler) createFacility(w http.ResponseWriter, r *http.Request) {
	var dto CreateFacilityDto

	if err := httputils.ReadJsonPayload(r, &dto); err != nil {
		httputils.SendJsonResponse(w, http.StatusBadRequest, httputils.ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	sub, _ := iam.ClaimsFromContext(r.Context()).GetSubject()
	resultChan := make(chan result)

	go func() {
		f, err := h.services.Facility.Create(r.Context(), domain.Facility{
			Code:      dto.Code,
			Name:      dto.Name,
			CreatedBy: sub,
		})
		resultChan <- result{f, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			httputils.SendJsonResponse(w, http.StatusInternalServerError, httputils.ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		httputils.SendJsonResponse(w, http.StatusCreated, httputils.ApiResponseDto{
			Message: "Sucessfully created facility!",
			Data:    res.data,
		})
	}
}

func (h *Handler) updateFacility(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	var dto UpdateFacilityDto

	if err := httputils.ReadJsonPayload(r, &dto); err != nil {
		httputils.SendJsonResponse(w, http.StatusBadRequest, httputils.ApiErrorResponseDto{
			Message:    err.Error(),
			Error:      "Bad Request",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	sub, _ := iam.ClaimsFromContext(r.Context()).GetSubject()
	resultChan := make(chan result)

	go func() {
		f, err := h.services.Facility.Update(r.Context(), domain.Facility{
			Code:      code,
			Name:      dto.Name,
			UpdatedBy: sub,
		})
		resultChan <- result{f, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			httputils.SendJsonResponse(w, http.StatusInternalServerError, httputils.ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		httputils.SendJsonResponse(w, http.StatusOK, httputils.ApiResponseDto{
			Message: "Sucessfully updated facility!",
			Data:    res.data,
		})
	}
}

func (h *Handler) removeFacility(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	resultChan := make(chan result)

	go func() {
		err := h.services.Facility.Remove(r.Context(), code)
		resultChan <- result{nil, err}
	}()

	select {
	case <-r.Context().Done():
		return
	case res := <-resultChan:
		if res.err != nil {
			httputils.SendJsonResponse(w, http.StatusInternalServerError, httputils.ApiErrorResponseDto{
				Message:    res.err.Error(),
				Error:      "Internal Server Error",
				StatusCode: http.StatusInternalServerError,
			})
			return
		}

		httputils.SendJsonResponse(w, http.StatusOK, httputils.ApiResponseDto{
			Message: "Successfully deleted facility!",
		})
	}

}
