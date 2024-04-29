package iam

import "context"

const ClaimsContextKey = "claims"

type Claims interface {
	GetSubject() (string, error)
	GetAuthorizationDetails() []string
}

func SetContext(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, ClaimsContextKey, claims)
}

func ClaimsFromContext(ctx context.Context) Claims {
	return ctx.Value(ClaimsContextKey).(Claims)
}
