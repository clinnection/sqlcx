// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package querytest

import (
	"time"
)

type Author struct {
	Name      string
	DeletedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Book struct {
	IsAmazing bool
	DeletedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
