package config

import (
	"time"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/kv"
)

type DBCleaner struct {
	Period time.Duration `fig:"period,required"`
}

func (c *config) DBCleaner() *DBCleaner {
	return c.dbCleaner.Do(func() interface{} {
		cfg := DBCleaner{}
		err := figure.
			Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "db_cleaner")).
			Please()
		if err != nil {
			panic(errors.WithMessage(err, "failed to figure out"))
		}

		return &cfg
	}).(*DBCleaner)
}
