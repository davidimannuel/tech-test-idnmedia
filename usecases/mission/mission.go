package mission

import (
	"context"
	"idnmedia/repositories"
	"idnmedia/repositories/mission"
	"idnmedia/usecases"
	"idnmedia/utils"
)

type usecase struct {
	missionRepo mission.Repository
}

func NewUsecase(missionRepo mission.Repository) Usecase {
	return &usecase{
		missionRepo: missionRepo,
	}
}

func (uc *usecase) Create(ctx context.Context, req *MissionEntity) (res MissionEntity, err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}
	res = *req
	res.Id, err = uc.missionRepo.Create(ctx, mission.MissionModel{
		Title:          req.Title,
		Description:    req.Description,
		GoldBounty:     req.GoldBounty,
		DeadlineSecond: req.DeadlineSecond,
		BaseModel: repositories.BaseModel{
			CreatedBy: playerCtx.Email,
			UpdatedBy: playerCtx.Email,
		},
	})
	if err != nil {
		return
	}

	return
}

func (uc *usecase) FindAllPagination(ctx context.Context, page, limit int) (res []MissionEntity, p usecases.Pagination, err error) {
	p = usecases.Pagination{
		CurrentPage: page,
		PerPage:     limit,
	}
	p.Validate()

	data, count, err := uc.missionRepo.FindAllPagination(ctx, p.GetOffset(), p.PerPage)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}
	p.TotalData = count

	for _, d := range data {
		res = append(res, MissionEntity{
			Id:             d.Id,
			Title:          d.Title,
			Description:    d.Description,
			GoldBounty:     d.GoldBounty,
			DeadlineSecond: d.DeadlineSecond,
			BaseEntity: usecases.BaseEntity{
				CreatedAt: d.CreatedAt,
				CreatedBy: d.CreatedBy,
				UpdatedAt: d.UpdatedAt,
				UpdatedBy: d.UpdatedBy,
			},
		})
	}

	return
}

func (uc *usecase) FindOneByID(ctx context.Context, id int) (res MissionEntity, err error) {

	data, err := uc.missionRepo.FindOneByID(ctx, id)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}

	res = MissionEntity{
		Id:             data.Id,
		Title:          data.Title,
		Description:    data.Description,
		GoldBounty:     data.GoldBounty,
		DeadlineSecond: data.DeadlineSecond,
		BaseEntity: usecases.BaseEntity{
			CreatedAt: data.CreatedAt,
			CreatedBy: data.CreatedBy,
			UpdatedAt: data.UpdatedAt,
			UpdatedBy: data.UpdatedBy,
		},
	}

	return
}
