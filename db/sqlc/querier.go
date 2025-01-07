// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountUser(ctx context.Context, dollar_1 sql.NullString) (int64, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error)
	GetMyProfile(ctx context.Context, username string) (User, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]GetUsersRow, error)
	SearchUser(ctx context.Context, arg SearchUserParams) ([]SearchUserRow, error)
	UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (UpdateAvatarRow, error)
	UpdatePosition(ctx context.Context, arg UpdatePositionParams) (UpdatePositionRow, error)
}

var _ Querier = (*Queries)(nil)
