// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type User struct {
	FirstName pgtype.Text `json:"firstName"`
	LastName  pgtype.Text `json:"lastName"`
	Age       pgtype.Int2 `json:"age"`
}
