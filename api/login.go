package api

import (
	"blog/common"
	"blog/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	parans := common.GetRequestJsonParam(r)

	userName := parans["username"].(string)
	passwd := parans["passwd"].(string)

	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)
}
