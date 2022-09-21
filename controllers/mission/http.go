package mission

import (
	"encoding/json"
	"net/http"
	"strconv"

	"idnmedia/controllers"
	"idnmedia/usecases/mission"
)

type httpController struct {
	controllers.BaseController
	missionUc mission.Usecase
}

func NewHttpController(missionUc mission.Usecase) *httpController {
	return &httpController{
		missionUc: missionUc,
	}
}

func (ctrl *httpController) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(MissionCreateRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		controllers.WriteResponse(w, http.StatusBadRequest, err.Error(), nil, nil)
	}

	result, err := ctrl.missionUc.Create(ctx, &mission.MissionEntity{
		Title:          req.Title,
		Description:    req.Description,
		GoldBounty:     req.GoldBounty,
		DeadlineSecond: req.DeadlineSecond,
	})
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	res := MissionResponse{
		Id:             result.Id,
		Title:          result.Title,
		Description:    result.Description,
		GoldBounty:     result.GoldBounty,
		DeadlineSecond: result.DeadlineSecond,
		CreatedAt:      result.CreatedAt,
		CreatedBy:      result.CreatedBy,
		UpdatedAt:      result.UpdatedAt,
		UpdatedBy:      result.UpdatedBy,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}

func (ctrl *httpController) FindAllPagination(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	result, pagination, err := ctrl.missionUc.FindAllPagination(ctx, page, limit)
	if err != nil {
		controllers.WriteResponse(w, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}
	res := []MissionResponse{}
	for _, v := range result {
		res = append(res, MissionResponse{
			Id:             v.Id,
			Title:          v.Title,
			Description:    v.Description,
			GoldBounty:     v.GoldBounty,
			DeadlineSecond: v.DeadlineSecond,
			CreatedAt:      v.CreatedAt,
			CreatedBy:      v.CreatedBy,
			UpdatedAt:      v.UpdatedAt,
			UpdatedBy:      v.UpdatedBy,
		})
	}
	p := controllers.Pagination{
		CurrentPage: pagination.CurrentPage,
		PerPage:     pagination.PerPage,
		LastPage:    pagination.GetLastPage(),
		TotalData:   pagination.TotalData,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, p)
}
