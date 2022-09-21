package player_mission

import (
	"context"
	"idnmedia/repositories"
	"time"
)

type Repository interface {
	Create(ctx context.Context, m PlayerMissionModel) (lastID int, err error)
	UpdateStatus(ctx context.Context, id int, status string) (err error)
	Delete(ctx context.Context, id int) (err error)
	FindAllByPlayerID(ctx context.Context, playerId int) (res []PlayerMissionModel, err error)
	FindOneByPlayerIDAndMissionID(ctx context.Context, playerId, missionId int) (res PlayerMissionModel, err error)
}

type PlayerMissionModel struct {
	Id           int
	PlayerId     int
	MissionId    int
	Status       string
	DeadlineTime *time.Time
	repositories.BaseModel

	MissionTitle       string
	MissionDescription string
	MissionGoldBounty  float64
}
