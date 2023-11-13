package pg

import (
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/rarimo/verify-proxy/internal/data"
)

const (
	verifyRequestsTableName = "verify_requests"
)

const (
	idColumnName        = "id"
	createdAtColumnName = "created_at"
)

type masterQ struct {
	db *pgdb.DB
}

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{
		db: db,
	}
}

func (q *masterQ) New() data.MasterQ {
	return NewMasterQ(q.db.Clone())
}

func (q *masterQ) VerifyRequestsQ() data.VerifyRequestsQ {
	return NewVerifyRequestsQ(q.db)
}

func (q *masterQ) Transaction(fn func() error) error {
	return q.db.Transaction(fn)
}
