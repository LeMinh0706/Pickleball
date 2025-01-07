// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const countUser = `-- name: CountUser :one
SELECT count(id)
FROM users
WHERE fullname ILIKE '%' || $1 || '%'
`

func (q *Queries) CountUser(ctx context.Context, dollar_1 sql.NullString) (int64, error) {
	row := q.db.QueryRowContext(ctx, countUser, dollar_1)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users(
    id,
    username,
    password,
    fullname,
    gender,
    avt,
    lat,
    lng
) VALUES(
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING fullname, gender, avt, lat, lng
`

type CreateUserParams struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Fullname string    `json:"fullname"`
	Gender   int32     `json:"gender"`
	Avt      string    `json:"avt"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
}

type CreateUserRow struct {
	Fullname string  `json:"fullname"`
	Gender   int32   `json:"gender"`
	Avt      string  `json:"avt"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Password,
		arg.Fullname,
		arg.Gender,
		arg.Avt,
		arg.Lat,
		arg.Lng,
	)
	var i CreateUserRow
	err := row.Scan(
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}

const getMyProfile = `-- name: GetMyProfile :one
SELECT id, username, password, fullname, gender, avt, lat, lng, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetMyProfile(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getMyProfile, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.Lat,
		&i.Lng,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password, fullname, gender, avt, lat, lng, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.Lat,
		&i.Lng,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :many
SELECT 
  id, 
  username, 
  fullname, 
  avt, 
  lat, 
  lng, 
  created_at,
  CAST(
    6371 * acos(
      cos(radians($1)) * cos(radians(lat)) * 
      cos(radians(lng) - radians($2)) + 
      sin(radians($1)) * sin(radians(lat))
    ) AS float8
  ) AS distance
FROM users
WHERE 
  (6371 * acos(
    cos(radians($1)) * cos(radians(lat)) * 
    cos(radians(lng) - radians($2)) + 
    sin(radians($1)) * sin(radians(lat))
  )) <= $3
ORDER BY distance
`

type GetUsersParams struct {
	Radians   float64 `json:"radians"`
	Radians_2 float64 `json:"radians_2"`
	Lat       float64 `json:"lat"`
}

type GetUsersRow struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Avt       string    `json:"avt"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
	Distance  float64   `json:"distance"`
}

func (q *Queries) GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getUsers, arg.Radians, arg.Radians_2, arg.Lat)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetUsersRow{}
	for rows.Next() {
		var i GetUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Fullname,
			&i.Avt,
			&i.Lat,
			&i.Lng,
			&i.CreatedAt,
			&i.Distance,
		); err != nil {
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

const searchUser = `-- name: SearchUser :many
SELECT id, fullname, gender, avt, lat, lng
FROM users
WHERE fullname ILIKE '%' || $1 || '%'
LIMIT $2
OFFSET $3
`

type SearchUserParams struct {
	Column1 sql.NullString `json:"column_1"`
	Limit   int32          `json:"limit"`
	Offset  int32          `json:"offset"`
}

type SearchUserRow struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Gender   int32     `json:"gender"`
	Avt      string    `json:"avt"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
}

func (q *Queries) SearchUser(ctx context.Context, arg SearchUserParams) ([]SearchUserRow, error) {
	rows, err := q.db.QueryContext(ctx, searchUser, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchUserRow{}
	for rows.Next() {
		var i SearchUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Fullname,
			&i.Gender,
			&i.Avt,
			&i.Lat,
			&i.Lng,
		); err != nil {
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

const updateAvatar = `-- name: UpdateAvatar :one
UPDATE users SET 
avt = $2
WHERE id = $1
RETURNING id, fullname, gender, avt, lat, lng
`

type UpdateAvatarParams struct {
	ID  uuid.UUID `json:"id"`
	Avt string    `json:"avt"`
}

type UpdateAvatarRow struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Gender   int32     `json:"gender"`
	Avt      string    `json:"avt"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (UpdateAvatarRow, error) {
	row := q.db.QueryRowContext(ctx, updateAvatar, arg.ID, arg.Avt)
	var i UpdateAvatarRow
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}

const updatePosition = `-- name: UpdatePosition :one
UPDATE users SET 
lat = $2, 
lng = $3
WHERE id = $1
RETURNING id, fullname, gender, avt, lat, lng
`

type UpdatePositionParams struct {
	ID  uuid.UUID `json:"id"`
	Lat float64   `json:"lat"`
	Lng float64   `json:"lng"`
}

type UpdatePositionRow struct {
	ID       uuid.UUID `json:"id"`
	Fullname string    `json:"fullname"`
	Gender   int32     `json:"gender"`
	Avt      string    `json:"avt"`
	Lat      float64   `json:"lat"`
	Lng      float64   `json:"lng"`
}

func (q *Queries) UpdatePosition(ctx context.Context, arg UpdatePositionParams) (UpdatePositionRow, error) {
	row := q.db.QueryRowContext(ctx, updatePosition, arg.ID, arg.Lat, arg.Lng)
	var i UpdatePositionRow
	err := row.Scan(
		&i.ID,
		&i.Fullname,
		&i.Gender,
		&i.Avt,
		&i.Lat,
		&i.Lng,
	)
	return i, err
}
