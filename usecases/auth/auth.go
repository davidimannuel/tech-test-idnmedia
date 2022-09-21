package auth

import (
	"context"
	"idnmedia/controllers/middlewares"
	"idnmedia/repositories/player"
	"idnmedia/utils"
)

type usecase struct {
	jwt        *middlewares.JWT
	playerRepo player.Repository
}

func NewUsecase(jwt *middlewares.JWT, playerRepo player.Repository) Usecase {
	return &usecase{
		jwt:        jwt,
		playerRepo: playerRepo,
	}
}

func (uc *usecase) Login(ctx context.Context, ent *AuthEntity) (token string, err error) {
	player, err := uc.playerRepo.FindOneByEmail(ctx, ent.Email)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get player by email")
		return
	}

	if !utils.BcryptCompare(player.Password, ent.Password) {
		utils.LogError(ctx, utils.FnTrace(), "password not match")
		err = utils.ErrInvalidPassword
		return
	}

	return uc.jwt.GenerateToken(ctx, player.Id, player.Email)
}

func (uc *usecase) Profile(ctx context.Context) (res AuthProfileEntity, err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}
	player, err := uc.playerRepo.FindOneByEmail(ctx, playerCtx.Email)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get player by email")
		return
	}

	res = AuthProfileEntity{
		Id:         player.Id,
		Name:       player.Name,
		Email:      player.Email,
		GoldAmount: player.GoldAmount,
	}

	return
}
