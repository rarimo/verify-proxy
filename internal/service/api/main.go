package api

import (
	"context"
	"net"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/config"
	"gitlab.com/rarimo/polygonid/verify-proxy/internal/service/core"
	dbcleaner "gitlab.com/rarimo/polygonid/verify-proxy/internal/service/core/db_cleaner"
)

type service struct {
	log         *logan.Entry
	copus       types.Copus
	listener    net.Listener
	verifyProxy core.VerifyProxy
	jwtCfg      *config.JWT
}

func newService(cfg config.Config) *service {
	return &service{
		log:         cfg.Log(),
		copus:       cfg.Copus(),
		listener:    cfg.Listener(),
		jwtCfg:      cfg.JWT(),
		verifyProxy: core.NewVerifyProxy(cfg),
	}
}

func Run(ctx context.Context, cfg config.Config) {
	svc := newService(cfg)

	go dbcleaner.NewVerifyRequestsCleaner(cfg).Run(ctx)

	svc.log.Info("Service started")
	defer svc.log.Info("Service stopped")
	ape.Serve(ctx, svc.router(), cfg, ape.ServeOpts{})
}
