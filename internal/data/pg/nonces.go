package pg

import (
	"database/sql"
	"github.com/kish1n/shortlink/internal/data"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const tableName = "links"

func newLinksQ(db *pgdb.DB) data.LinksQ {
	return &nonceQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type nonceQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *nonceQ) Get() (*data.CoupleLinks, error) {
	var result data.CoupleLinks
	err := q.db.Get(&result, q.sql.Select("*").From(tableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get link from db")
	}
	return &result, nil
}

func (q *nonceQ) Select() ([]data.CoupleLinks, error) {
	var result []data.CoupleLinks
	err := q.db.Select(&result, q.sql.Select("*").From(tableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select links from db")
	}
	return result, nil
}

func (q *nonceQ) Insert(value data.CoupleLinks) (*data.CoupleLinks, error) {
	clauses := structs.Map(value)

	var result data.CoupleLinks
	stmt := sq.Insert(tableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert link to db")
	}
	return &result, nil
}

func (q *nonceQ) Update(value data.CoupleLinks) (*data.CoupleLinks, error) {
	clauses := structs.Map(value)

	var result data.CoupleLinks
	stmt := q.sql.Update(tableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update link in db")
	}
	return &result, nil
}

func (q *nonceQ) Delete() error {
	err := q.db.Exec(q.sql.Delete(tableName))
	if err != nil {
		return errors.Wrap(err, "failed to delete links from db")
	}
	return nil
}

func (q *nonceQ) FilterByAddress(addresses ...string) data.LinksQ {
	pred := sq.Eq{"address": addresses}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *nonceQ) FilterExpired() data.LinksQ {
	q.sql = sq.StatementBuilder.Where("expiresat < ?", time.Now().Unix())
	return q
}
