package views

import (
	"blog/common"
	"blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {

	categoryTemplate := common.Template.Category

	//获取页面路径
	path := r.URL.Path
	//截取出分类ID
	cidStr := strings.TrimPrefix(path, "/c/")
	cId, err := strconv.Atoi(cidStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("路径不匹配"))
	}

	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		//转为int
		page, _ = strconv.Atoi(pageStr)
	}

	limit := 10
	limitStr := r.Form.Get("limit")
	if limitStr != "" {
		limit, _ = strconv.Atoi(limitStr)
	}

	hr, err := service.GetCategoryIndexInfo(cId, page, limit)

	if err != nil {
		log.Println("Category获取数据出错：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员"))
	}

	categoryTemplate.WriteData(w, hr)
}
