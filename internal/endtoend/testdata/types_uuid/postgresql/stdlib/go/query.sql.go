// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/google/uuid"
)

const find = `-- name: Find :one
SELECT bar FROM foo WHERE baz = $1
`

func (q *Queries) Find(ctx context.Context, baz uuid.UUID) (uuid.NullUUID, error) {
	row := q.db.QueryRowContext(ctx, find, baz)
	var bar uuid.NullUUID
	err := row.Scan(&bar)
	return bar, err
}

const list = `-- name: List :many
SELECT description, bar, baz FROM foo
`

func (q *Queries) List(ctx context.Context) ([]Foo, error) {
	rows, err := q.db.QueryContext(ctx, list)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.Description, &i.Bar, &i.Baz); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
