package auth

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"strings"

	httputils "github.com/cmclaughlin24/soteria-be/pkg/http-utils"
)

type Authenticator func(r *http.Request) (Claims, error)

func Authenticate(authenitactors ...Authenticator) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, authenticator := range authenitactors {
				claims, err := authenticator(r)

				if err != nil {
					// Fixme: Add log message indicating the authenciation method failed.
					continue
				}

				valOfClaims := reflect.ValueOf(claims)

				if valOfClaims.IsValid() && !valOfClaims.IsZero() {
					r = r.WithContext(SetContext(r.Context(), claims))
					next.ServeHTTP(w, r)
					return
				}
			}

			httputils.SendJsonResponse(w, http.StatusUnauthorized, httputils.ApiErrorResponseDto{
				Message:    "Unauthorized",
				Error:      "Unauthorized",
				StatusCode: http.StatusUnauthorized,
			})
		})
	}
}

func Authorize(resource, action string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims := ClaimsFromContext(r.Context())
			permissions := UnpackPermissions(claims.GetAuthorizationDetails())

			for res, actions := range permissions {
				if res != resource {
					continue
				}

				for _, act := range actions {
					if act == action {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			httputils.SendJsonResponse(w, http.StatusForbidden, httputils.ApiErrorResponseDto{
				Message:    "Forbidden",
				Error:      "Forbidden",
				StatusCode: http.StatusForbidden,
			})
		})
	}
}

type AccessTokenVerifier[T Claims] interface {
	VerifyAccessToken(context.Context, string) (T, error)
}

func AuthenticateAccessToken[T Claims](verifier AccessTokenVerifier[T]) Authenticator {
	return func(r *http.Request) (Claims, error) {
		token := strings.Split(r.Header.Get("Authorization"), " ")

		if len(token) != 2 {
			return nil, errors.New("authorization header is not a tuple with type and token")
		}

		if token[0] != "Bearer" {
			return nil, errors.New("authorization header is not type \"Bearer\"")
		}

		return verifier.VerifyAccessToken(r.Context(), token[1])
	}
}

type ApiKeyVerifier[T Claims] interface {
	VerifyApiKey(context.Context, string) (T, error)
}

func AuthenticateApiKey[T Claims](verifier ApiKeyVerifier[T]) Authenticator {
	return func(r *http.Request) (Claims, error) {
		key := strings.Split(r.Header.Get("Authorization"), " ")

		if len(key) != 2 {
			return nil, errors.New("authorization header is not a tuple with type and token")
		}

		if key[0] != "ApiKey" {
			return nil, errors.New("authorization header is not type \"ApiKey\"")
		}

		return verifier.VerifyApiKey(r.Context(), key[1])
	}
}
