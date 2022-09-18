package views

import (
	"blog/common"
	"blog/service"
	"net/http"
)

func (*HTMLApi) Pigeonhole(w http.ResponseWriter, r *http.Request) {
	pigeonhole := common.Template.Pigeonhole
	hr, _ := service.FindPostPigeonhole()
	pigeonhole.WriteData(w, hr)
}
