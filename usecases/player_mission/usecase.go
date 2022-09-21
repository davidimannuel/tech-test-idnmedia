package player_mission

import (
	"context"
	"idnmedia/usecases"
)

type Usecase interface {
	FindAll(ctx context.Context) (res []PlayerMissionEntity, err error)
	Assign(ctx context.Context, missionID int) (res PlayerMissionEntity, err error)
	Progress(ctx context.Context, missionID int) (err error)
	Complete(ctx context.Context, missionID int) (res PlayerMissionEntity, err error)
}

type PlayerMissionEntity struct {
	Id        int
	PlayerId  int
	MissionId int
	Status    string
	usecases.BaseEntity

	MissionTitle       string
	MissionDescription string
	MissionGoldBounty  float64
}
