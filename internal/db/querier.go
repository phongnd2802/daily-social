// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package db

import (
	"context"
)

type Querier interface {
	CheckUserBaseExists(ctx context.Context, userEmail string) (int64, error)
	CreateUserBase(ctx context.Context, arg CreateUserBaseParams) (UserBase, error)
	CreateUserProfile(ctx context.Context, arg CreateUserProfileParams) (UserProfile, error)
	CreateUserSession(ctx context.Context, arg CreateUserSessionParams) (UserSession, error)
	GetUserBaseByEmail(ctx context.Context, userEmail string) (UserBase, error)
	UpdateUserVerify(ctx context.Context, userHash string) (UserBase, error)
}

var _ Querier = (*Queries)(nil)
