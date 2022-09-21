package controllers

import (
	"encoding/json"
	"idnmedia/utils"
	"net/http"
)

var (
	httpCodeError = map[error]int{
		utils.ErrInvalidPassword:      http.StatusUnauthorized,
		utils.ErrInvalidPlayerSession: http.StatusUnauthorized,
		utils.ErrDataNotFound:         http.StatusNotFound,
	}
)

type BaseController struct{}

func (BaseController) GetErrHTTPCode(err error) int {
	v, exist := httpCodeError[err]
	if !exist {
		return http.StatusInternalServerError
	}
	return v
}

type BaseResponse struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

type Pagination struct {
	CurrentPage int `json:"currentPage"`
	PerPage     int `json:"perPage"`
	LastPage    int `json:"lastPage"`
	TotalData   int `json:"totalData"`
}

func WriteResponse(w http.ResponseWriter, statusCode int, message string, data, meta interface{}) {
	bRes, _ := json.Marshal(BaseResponse{
		Data: data,
		Meta: meta,
		Msg:  message,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(bRes)
}
