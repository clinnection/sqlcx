// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package querytest

import (
	"database/sql"
	"time"
)

type Author struct {
	ID       int64
	Name     string
	ParentID sql.NullInt64
}

type City struct {
	CityID  int64
	MayorID int64
}

type Mayor struct {
	MayorID  int64
	FullName string
}

type Medium struct {
	MediaID        int64
	MediaCreatedAt time.Time
	MediaHash      string
	MediaDirectory string
	MediaAuthorID  int64
	MediaWidth     int64
	MediaHeight    int64
}

type SuperAuthor struct {
	SuperID       int64
	SuperName     string
	SuperParentID sql.NullInt64
}

type User struct {
	UserID int64
	CityID sql.NullInt64
}

type Users2 struct {
	UserID          int64
	UserNickname    string
	UserEmail       string
	UserDisplayName string
	UserPassword    sql.NullString
	UserGoogleID    sql.NullString
	UserAppleID     sql.NullString
	UserBio         string
	UserCreatedAt   time.Time
	UserAvatarID    sql.NullInt64
}
