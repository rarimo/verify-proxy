package middleware

import "github.com/pkg/errors"

const (
	AuthorizationHeaderName = "Authorization"
	BearerTokenPrefix       = "Bearer "
)

var (
	ErrInvalidAccessToken = errors.New("invalid access token")
	ErrTokenHasExpired    = errors.New("token has expired")
)
