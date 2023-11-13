package dbcleaner

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/running"

	"github.com/rarimo/verify-proxy/internal/config"
	"github.com/rarimo/verify-proxy/internal/data"
	"github.com/rarimo/verify-proxy/internal/data/pg"
)

type Cleaner interface {
	Run(ctx context.Context)
}

type verifRequestsCleaner struct {
	log            *logan.Entry
	db             data.MasterQ
	cleanPeriod    time.Duration
	entityLifeTime time.Duration
}

func NewVerifyRequestsCleaner(cfg config.Config) Cleaner {
	return &verifRequestsCleaner{
		log:            cfg.Log(),
		db:             pg.NewMasterQ(cfg.DB()),
		entityLifeTime: cfg.JWT().ExpirationTime,
		cleanPeriod:    cfg.DBCleaner().Period,
	}
}

func (c *verifRequestsCleaner) Run(ctx context.Context) {
	running.WithBackOff(
		ctx,
		c.log,
		dbCleanerServiceName,
		c.removeOutdatedVerifyRequests,
		c.cleanPeriod, c.cleanPeriod, c.cleanPeriod,
	)
}

func (c *verifRequestsCleaner) removeOutdatedVerifyRequests(_ context.Context) error {
	err := c.db.VerifyRequestsQ().WhereCreatedAtLt(time.Now().UTC().Add(-c.entityLifeTime)).Delete()
	if err != nil {
		return errors.Wrap(err, "failed to delete expired verification requests")
	}

	return nil
}
