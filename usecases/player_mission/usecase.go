package player_mission

import (
	"context"
	"idnmedia/usecases"
	"time"
)

type Usecase interface {
	FindAll(ctx context.Context) (res []PlayerMissionEntity, err error)
	Assign(ctx context.Context, missionID int) (res PlayerMissionEntity, err error)
	Complete(ctx context.Context, missionID int) (res PlayerMissionEntity, err error)
}

type PlayerMissionEntity struct {
	Id           int
	PlayerId     int
	MissionId    int
	Status       string
	DeadlineTime *time.Time
	usecases.BaseEntity

	MissionTitle       string
	MissionDescription string
	MissionGoldBounty  float64
}
