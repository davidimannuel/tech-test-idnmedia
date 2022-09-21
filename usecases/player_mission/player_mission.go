package player_mission

import (
	"context"
	"database/sql"
	"idnmedia/constants"
	"idnmedia/repositories"
	"idnmedia/repositories/mission"
	player "idnmedia/repositories/player"
	playerMission "idnmedia/repositories/player_mission"
	"idnmedia/usecases"
	"idnmedia/utils"
)

type usecase struct {
	missionRepo       mission.Repository
	playerMissionRepo playerMission.Repository
	playerRepo        player.Repository
	dbTx              repositories.DBTransactional
}

func NewUsecase(missionRepo mission.Repository, playerMissionRepo playerMission.Repository, playerRepo player.Repository, dbTx repositories.DBTransactional) Usecase {
	return &usecase{
		missionRepo:       missionRepo,
		playerMissionRepo: playerMissionRepo,
		playerRepo:        playerRepo,
		dbTx:              dbTx,
	}
}

func (uc *usecase) FindAll(ctx context.Context) (res []PlayerMissionEntity, err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}

	data, err := uc.playerMissionRepo.FindAllByPlayerID(ctx, playerCtx.Id)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get player missions")
		return
	}

	for _, d := range data {
		res = append(res, PlayerMissionEntity{
			Id:        d.Id,
			PlayerId:  d.PlayerId,
			MissionId: d.MissionId,
			Status:    d.Status,
			BaseEntity: usecases.BaseEntity{
				CreatedAt: d.CreatedAt,
				CreatedBy: d.CreatedBy,
				UpdatedAt: d.UpdatedAt,
				UpdatedBy: d.UpdatedBy,
			},
			MissionTitle:       d.MissionTitle,
			MissionDescription: d.MissionDescription,
			MissionGoldBounty:  d.MissionGoldBounty,
		})
	}

	return
}

func (uc *usecase) Assign(ctx context.Context, missionID int) (res PlayerMissionEntity, err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}

	mission, err := uc.missionRepo.FindOneByID(ctx, missionID)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get mission")
		if err == sql.ErrNoRows {
			err = utils.ErrDataNotFound
		}
		return
	}

	pm, err := uc.playerMissionRepo.FindOneByPlayerIDAndMissionID(ctx, playerCtx.Id, missionID)
	if err != nil && err != sql.ErrNoRows {
		utils.LogError(ctx, utils.FnTrace(), "error get player mission")
		return
	}

	if pm.Id > 0 {
		err = utils.ErrMissionAlreadyAssigned
		utils.LogError(ctx, utils.FnTrace(), "error mission already assigned")
		return
	}

	lastId, err := uc.playerMissionRepo.Create(ctx, playerMission.PlayerMissionModel{
		PlayerId:  playerCtx.Id,
		MissionId: missionID,
		Status:    constants.PLAYER_MISSION_STATUS_PENDING,
		BaseModel: repositories.BaseModel{
			CreatedBy: playerCtx.Email,
			UpdatedBy: playerCtx.Email,
		},
	})
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error create player mission")
		return
	}

	res = PlayerMissionEntity{
		Id:                 lastId,
		PlayerId:           playerCtx.Id,
		MissionId:          missionID,
		Status:             constants.PLAYER_MISSION_STATUS_PENDING,
		MissionTitle:       mission.Title,
		MissionDescription: mission.Description,
	}

	return
}

func (uc *usecase) Progress(ctx context.Context, missionID int) (err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}

	playerMission, err := uc.playerMissionRepo.FindOneByPlayerIDAndMissionID(ctx, playerCtx.Id, missionID)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get player mission")
		if err == sql.ErrNoRows {
			err = utils.ErrDataNotFound
		}
		return
	}

	if playerMission.Status != constants.PLAYER_MISSION_STATUS_PENDING {
		err = utils.ErrInvalidMissionStatus
		utils.LogError(ctx, utils.FnTrace(), "invalid mission status %v", playerMission.Status)
		return
	}

	err = uc.playerMissionRepo.UpdateStatus(ctx, playerMission.Id, constants.PLAYER_MISSION_STATUS_IN_PROGESS)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error udpate status player mission")
		return
	}

	return
}

func (uc *usecase) Complete(ctx context.Context, missionID int) (res PlayerMissionEntity, err error) {
	playerCtx, err := utils.GetCtxPLayer(ctx)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get ctx player")
		return
	}

	playerMission, err := uc.playerMissionRepo.FindOneByPlayerIDAndMissionID(ctx, playerCtx.Id, missionID)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error get player mission")
		if err == sql.ErrNoRows {
			err = utils.ErrDataNotFound
		}
		return
	}

	if playerMission.Status != constants.PLAYER_MISSION_STATUS_IN_PROGESS {
		err = utils.ErrInvalidMissionStatus
		utils.LogError(ctx, utils.FnTrace(), "invalid mission status %v", playerMission.Status)
		return
	}

	// open transactions
	tx, err := uc.dbTx.Begin(ctx, &sql.TxOptions{})
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error open tx")
		return
	}
	defer tx.Rollback()

	ctx = utils.SetCtxDBTx(ctx, tx)

	err = uc.playerMissionRepo.UpdateStatus(ctx, playerMission.Id, constants.PLAYER_MISSION_STATUS_COMPLETED)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error udpate status player mission")
		return
	}

	_, err = uc.playerRepo.AddGoldAmountByPlayerId(ctx, playerCtx.Id, playerMission.MissionGoldBounty)
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error udpate status player mission")
		return
	}

	err = tx.Commit()
	if err != nil {
		utils.LogError(ctx, utils.FnTrace(), "error commit")
		return
	}

	res = PlayerMissionEntity{
		Status:             constants.PLAYER_MISSION_STATUS_COMPLETED,
		MissionTitle:       playerMission.MissionTitle,
		MissionDescription: playerMission.MissionDescription,
		MissionGoldBounty:  playerMission.MissionGoldBounty,
	}

	return
}
