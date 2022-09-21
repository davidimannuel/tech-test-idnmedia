package player

import (
	"context"
	"idnmedia/repositories"
)

type Repository interface {
	FindOneByEmail(ctx context.Context, username string) (row PlayerModel, err error)
	AddGoldAmountByPlayerId(ctx context.Context, id int, goldAmount float64) (currentGold float64, err error)
}

type PlayerModel struct {
	Id         int
	Name       string
	Email      string
	Password   string
	GoldAmount float64
	repositories.BaseModel
}
