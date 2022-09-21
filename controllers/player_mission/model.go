package player_mission

import "time"

type MissionCreateRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	GoldBounty  float64 `json:"goldBounty"`
}

type PlayerMissionResponse struct {
	MissionId          int        `json:"missionId"`
	Status             string     `json:"status"`
	MissionTitle       string     `json:"missionTitle"`
	MissionDescription string     `json:"missionDescription"`
	MissionGoldBounty  float64    `json:"missionGoldBounty"`
	DeadlineTime       *time.Time `json:"deadlineTime"`
	CreatedAt          time.Time  `json:"createdAt"`
	CreatedBy          string     `json:"createdBy"`
	UpdatedAt          time.Time  `json:"updatedAt"`
	UpdatedBy          string     `json:"updatedBy"`
}
