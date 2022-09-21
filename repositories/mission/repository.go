package mission

import (
	"context"
	"idnmedia/repositories"
)

type Repository interface {
	Create(ctx context.Context, m MissionModel) (lastID int, err error)
	Edit(ctx context.Context, m MissionModel) (err error)
	FindAll(ctx context.Context) (res []MissionModel, err error)
	FindAllPagination(ctx context.Context, offset, limit int) (res []MissionModel, count int, err error)
	FindOneByID(ctx context.Context, id int) (res MissionModel, err error)
}

type MissionModel struct {
	Id          int
	Title       string
	Description string
	GoldBounty  float64
	repositories.BaseModel
}
