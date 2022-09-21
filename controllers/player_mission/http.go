package player_mission

import (
	"net/http"
	"strconv"

	"idnmedia/controllers"
	playerMission "idnmedia/usecases/player_mission"

	"github.com/gorilla/mux"
)

type httpController struct {
	controllers.BaseController
	playerMissionUc playerMission.Usecase
}

func NewHttpController(playerMissionUc playerMission.Usecase) *httpController {
	return &httpController{
		playerMissionUc: playerMissionUc,
	}
}

func (ctrl *httpController) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	result, err := ctrl.playerMissionUc.FindAll(ctx)
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	res := []PlayerMissionResponse{}
	for _, v := range result {
		res = append(res, PlayerMissionResponse{
			MissionId:          v.MissionId,
			MissionTitle:       v.MissionTitle,
			MissionDescription: v.MissionDescription,
			MissionGoldBounty:  v.MissionGoldBounty,
			Status:             v.Status,
			DeadlineTime:       v.DeadlineTime,
			CreatedAt:          v.CreatedAt,
			CreatedBy:          v.CreatedBy,
			UpdatedAt:          v.UpdatedAt,
			UpdatedBy:          v.UpdatedBy,
		})
	}

	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}

func (ctrl *httpController) Assign(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	missionId, err := strconv.Atoi(mux.Vars(r)["missionId"])
	if err != nil {
		controllers.WriteResponse(w, http.StatusBadRequest, err.Error(), nil, nil)
	}

	result, err := ctrl.playerMissionUc.Assign(ctx, missionId)
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	res := PlayerMissionResponse{
		MissionId:          result.MissionId,
		Status:             result.Status,
		DeadlineTime:       result.DeadlineTime,
		MissionTitle:       result.MissionTitle,
		MissionDescription: result.MissionDescription,
		MissionGoldBounty:  result.MissionGoldBounty,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}

func (ctrl *httpController) Complete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	missionId, err := strconv.Atoi(mux.Vars(r)["missionId"])
	if err != nil {
		controllers.WriteResponse(w, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	result, err := ctrl.playerMissionUc.Complete(ctx, missionId)
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	res := PlayerMissionResponse{
		MissionId:          result.MissionId,
		Status:             result.Status,
		MissionTitle:       result.MissionTitle,
		MissionDescription: result.MissionDescription,
		MissionGoldBounty:  result.MissionGoldBounty,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}
