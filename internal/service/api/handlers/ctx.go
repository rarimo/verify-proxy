package handlers

import (
	"context"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/core"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	verifyProxyCtxKey
	JWTBasicClaimsCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxVerifyProxy(verifyProxy core.VerifyProxy) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, verifyProxyCtxKey, verifyProxy)
	}
}

func VerifyProxy(r *http.Request) core.VerifyProxy {
	return r.Context().Value(verifyProxyCtxKey).(core.VerifyProxy)
}

func JWTBasicClaims(r *http.Request) *core.BasicJWTClaims {
	return r.Context().Value(JWTBasicClaimsCtxKey).(*core.BasicJWTClaims)
}
