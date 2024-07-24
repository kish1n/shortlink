package pg

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/kish1n/shortlink/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

const tableName = "links"

func newLinksQ(db *pgdb.DB) data.LinksQ {
	return &LinksQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type LinksQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *LinksQ) FilterByOriginal(original string) (data.CoupleLinks, error) {
	builder := data.NewCoupleLinksBuilder()
	result := builder.
		SetShortened("").
		SetOriginal(original).
		Build()

	stmt := sq.Select("*").From("links").Where(sq.Eq{"original": original})
	err := q.db.Get(&result, stmt)
	if err != nil {
		return result, errors.Wrap(err, "failed to select by origin link in db")
	}

	return result, nil
}

func (q *LinksQ) Insert(value data.CoupleLinks) (*data.CoupleLinks, error) {
	clauses := structs.Map(value)

	var result data.CoupleLinks
	stmt := sq.Insert(tableName).SetMap(clauses).Suffix("returning *")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert link to db")
	}
	return &result, nil
}
