package auth

import (
	"encoding/json"
	"net/http"

	"idnmedia/controllers"
	"idnmedia/usecases/auth"
)

type httpController struct {
	controllers.BaseController
	authUC auth.Usecase
}

func NewHttpController(authUC auth.Usecase) *httpController {
	return &httpController{
		authUC: authUC,
	}
}

func (ctrl *httpController) Login(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req := new(LoginRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		controllers.WriteResponse(w, http.StatusBadRequest, err.Error(), nil, nil)
	}

	result, err := ctrl.authUC.Login(ctx, &auth.AuthEntity{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		controllers.WriteResponse(w, ctrl.GetErrHTTPCode(err), err.Error(), nil, nil)
		return
	}
	res := LoginResponse{
		Token: result,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}

func (ctrl *httpController) Profile(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	result, err := ctrl.authUC.Profile(ctx)
	if err != nil {
		controllers.WriteResponse(w, ctrl.GetErrHTTPCode(err), err.Error(), nil, nil)
		return
	}
	res := ProfileResponse{
		Id:         result.Id,
		Name:       result.Name,
		Email:      result.Email,
		GoldAmount: result.GoldAmount,
	}
	controllers.WriteResponse(w, http.StatusOK, "", res, nil)
}
