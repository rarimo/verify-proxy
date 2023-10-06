package pg

import (
	"database/sql"
	"time"

	"github.com/fatih/structs"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	sq "github.com/Masterminds/squirrel"

	"gitlab.com/rarimo/polygonid/verify-proxy/internal/data"
)

type verifyRequestsQ struct {
	db  *pgdb.DB
	sel sq.SelectBuilder
	del sq.DeleteBuilder
}

func NewVerifyRequestsQ(db *pgdb.DB) data.VerifyRequestsQ {
	return &verifyRequestsQ{
		db:  db,
		sel: sq.Select("*").From(verifyRequestsTableName),
		del: sq.Delete(verifyRequestsTableName),
	}
}

func (q *verifyRequestsQ) New() data.VerifyRequestsQ {
	return NewVerifyRequestsQ(q.db.Clone())
}

func (q *verifyRequestsQ) Select() ([]data.VerifyRequest, error) {
	var result []data.VerifyRequest

	err := q.db.Select(&result, q.sel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to select rows")
	}

	return result, nil
}

func (q *verifyRequestsQ) Get() (*data.VerifyRequest, error) {
	var result data.VerifyRequest

	err := q.db.Get(&result, q.sel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, errors.Wrap(err, "failed to get row")
	}

	return &result, nil
}

func (q *verifyRequestsQ) Insert(verifyRequest *data.VerifyRequest) error {
	err := q.db.Exec(
		sq.Insert(verifyRequestsTableName).
			SetMap(structs.Map(verifyRequest)),
	)
	if err != nil {
		return errors.Wrap(err, "failed to insert rows")
	}

	return nil
}

func (q *verifyRequestsQ) Update(verifyRequest *data.VerifyRequest) error {
	err := q.db.Exec(
		sq.Update(verifyRequestsTableName).
			SetMap(structs.Map(verifyRequest)).
			Where(sq.Eq{idColumnName: verifyRequest.ID}),
	)
	if err != nil {
		return errors.Wrap(err, "failed to update rows")
	}

	return nil
}

func (q *verifyRequestsQ) Delete() error {
	err := q.db.Exec(q.del)
	if err != nil {
		return errors.Wrap(err, "failed to delete rows")
	}

	return nil
}

func (q *verifyRequestsQ) WhereID(id uuid.UUID) data.VerifyRequestsQ {
	q.sel = q.sel.Where(sq.Eq{idColumnName: id})
	q.del = q.del.Where(sq.Eq{idColumnName: id})
	return q
}

func (q *verifyRequestsQ) WhereCreatedAtLt(createdAt time.Time) data.VerifyRequestsQ {
	q.sel = q.sel.Where(sq.Lt{createdAtColumnName: &createdAt})
	q.del = q.del.Where(sq.Lt{createdAtColumnName: &createdAt})
	return q
}
