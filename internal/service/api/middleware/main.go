package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/structs"
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/config"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/api/handlers"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/core"
)

func AuthMiddleware(cfg *config.JWT) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := GetBearer(r)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			jwtClaim, err := ValidateBearer(cfg, token)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), handlers.JWTBasicClaimsCtxKey, jwtClaim)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ValidateBearer(cfg *config.JWT, tokenRaw string) (*core.BasicJWTClaims, error) {
	claims := jwt.MapClaims(structs.Map(core.BasicJWTClaims{}))

	_, err := jwt.ParseWithClaims(tokenRaw, claims, func(token *jwt.Token) (interface{}, error) {
		return cfg.SecretKey, nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse token")
	}

	claimsMarshaled, err := json.Marshal(claims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal claims")
	}

	var parsedClaims core.BasicJWTClaims
	err = json.Unmarshal(claimsMarshaled, &parsedClaims)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal claims")
	}

	if parsedClaims.Exp < time.Now().UTC().Add(-cfg.ExpirationTime).Unix() {
		return nil, ErrTokenHasExpired
	}

	return &parsedClaims, nil
}

func GetBearer(r *http.Request) (string, error) {
	authHeader := r.Header.Get(AuthorizationHeaderName)
	authHeaderSplit := strings.Split(authHeader, BearerTokenPrefix)

	if len(authHeaderSplit) != 2 {
		return "", ErrInvalidAccessToken
	}

	return authHeaderSplit[1], nil
}
