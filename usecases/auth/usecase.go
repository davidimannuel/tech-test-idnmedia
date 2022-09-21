package auth

import (
	"context"
)

type Usecase interface {
	Login(ctx context.Context, ent *AuthEntity) (token string, err error)
	Profile(ctx context.Context) (res AuthProfileEntity, err error)
}

type AuthEntity struct {
	Email    string
	Password string
}

type AuthProfileEntity struct {
	Id         int
	Name       string
	Email      string
	GoldAmount float64
}
