package core

import (
	"time"

	"github.com/google/uuid"

	"github.com/rarimo/verify-proxy/internal/config"
	"github.com/rarimo/verify-proxy/internal/data"
	"github.com/rarimo/verify-proxy/internal/data/pg"
	"github.com/rarimo/verify-proxy/internal/service/api/requests"
)

type VerifyProxy interface {
	NewVerificationRequest() (*uuid.UUID, string, error)
	VerifyCallback(*requests.VerificationCallbackRequest) error
	GetJWZToken(uuid.UUID) (string, error)
}

type verifyProxy struct {
	db        data.MasterQ
	jwtSecret []byte
	jwtExp    time.Duration
}

func NewVerifyProxy(cfg config.Config) VerifyProxy {
	return &verifyProxy{
		db:        pg.NewMasterQ(cfg.DB()),
		jwtSecret: cfg.JWT().SecretKey,
		jwtExp:    cfg.JWT().ExpirationTime,
	}
}
