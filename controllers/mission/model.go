package mission

import "time"

type MissionCreateRequest struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	GoldBounty     float64 `json:"goldBounty"`
	DeadlineSecond int     `json:"deadlineSecond"`
}

type MissionResponse struct {
	Id             int       `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	GoldBounty     float64   `json:"goldBounty"`
	DeadlineSecond int       `json:"deadlineSecond"`
	CreatedAt      time.Time `json:"createdAt"`
	CreatedBy      string    `json:"createdBy"`
	UpdatedAt      time.Time `json:"updatedAt"`
	UpdatedBy      string    `json:"updatedBy"`
}
