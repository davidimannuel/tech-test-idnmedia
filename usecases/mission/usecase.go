package mission

import (
	"context"
	"idnmedia/usecases"
)

type Usecase interface {
	Create(ctx context.Context, req *MissionEntity) (res MissionEntity, err error)
	FindAllPagination(ctx context.Context, page, limit int) (res []MissionEntity, p usecases.Pagination, err error)
	FindOneByID(ctx context.Context, id int) (res MissionEntity, err error)
}

type MissionEntity struct {
	Id          int
	Title       string
	Description string
	GoldBounty  float64
	usecases.BaseEntity
}
