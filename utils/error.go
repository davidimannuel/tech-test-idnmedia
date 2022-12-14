package utils

import "errors"

var (
	ErrInvalidPassword        = errors.New("invalid_password")
	ErrInvalidPlayerSession   = errors.New("invalid_player_session")
	ErrDataNotFound           = errors.New("data_not_found")
	ErrMissionAlreadyAssigned = errors.New("mission_already_assigned")
	ErrInvalidMissionStatus   = errors.New("invalid_mission_status")
	ErrExpiredMission         = errors.New("expired_mission_status")
)
