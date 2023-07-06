// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: query.sql

package querytest

import (
	"context"
)

const deleteReadyWithCTE = `-- name: DeleteReadyWithCTE :many
WITH ready_ids AS (
	SELECT id FROM bar WHERE ready
)
DELETE FROM bar WHERE id IN (SELECT id FROM ready_ids)
RETURNING id
`

func (q *Queries) DeleteReadyWithCTE(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, deleteReadyWithCTE)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var id int32
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		items = append(items, id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
