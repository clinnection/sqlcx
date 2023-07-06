// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: batch.go

package querytest

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrBatchAlreadyClosed = errors.New("batch already closed")
)

const insertValues = `-- name: InsertValues :batchone
INSERT INTO foo (a, b)
VALUES ($1, $2)
ON CONFLICT DO NOTHING
RETURNING a, b
`

type InsertValuesBatchResults struct {
	br     pgx.BatchResults
	tot    int
	closed bool
}

type InsertValuesParams struct {
	A pgtype.Int4
	B pgtype.Int4
}

func (q *Queries) InsertValues(ctx context.Context, arg []*InsertValuesParams) *InsertValuesBatchResults {
	batch := &pgx.Batch{}
	for _, a := range arg {
		vals := []interface{}{
			a.A,
			a.B,
		}
		batch.Queue(insertValues, vals...)
	}
	br := q.db.SendBatch(ctx, batch)
	return &InsertValuesBatchResults{br, len(arg), false}
}

func (b *InsertValuesBatchResults) QueryRow(f func(int, *Foo, error)) {
	defer b.br.Close()
	for t := 0; t < b.tot; t++ {
		var i Foo
		if b.closed {
			if f != nil {
				f(t, nil, ErrBatchAlreadyClosed)
			}
			continue
		}
		row := b.br.QueryRow()
		err := row.Scan(&i.A, &i.B)
		if f != nil {
			f(t, &i, err)
		}
	}
}

func (b *InsertValuesBatchResults) Close() error {
	b.closed = true
	return b.br.Close()
}
