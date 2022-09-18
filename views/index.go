package views

import (
	"blog/common"
	"blog/service"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Date(layout string) string {
	return time.Now().Format(layout)
}

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(strs []string, index int) string {
	return strs[index+1]
}

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {

	index := common.Template.Index

	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
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

	path := r.URL.Path
	slug := strings.TrimPrefix(path, "/")

	hr, err := service.GetAllIndexInfo(slug, page, limit)

	if err != nil {
		log.Println("Index获取数据出错：", err)
		index.WriteError(w, errors.New("系统错误，请联系管理员!!"))
	}

	index.WriteData(w, hr)
}

func (*HTMLApi) Details(w http.ResponseWriter, r *http.Request) {

	details := common.Template.Detail

	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/p/")
	//7.html
	pIdStr = strings.TrimSuffix(pIdStr, ".html")

	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		details.WriteError(w, errors.New("路径不匹配"))
	}

	hr, err := service.GetPostDetails(pId)
	if err != nil {
		log.Println("Details获取数据出错：", err)
		details.WriteError(w, errors.New("系统错误，请联系管理员!!"))
	}

	details.WriteData(w, hr)

}

func (*HTMLApi) Writing(w http.ResponseWriter, r *http.Request) {

	writing := common.Template.Writing

	hr, err := service.GetWriting()
	if err != nil {
		log.Println("Writing获取数据出错：", err)
		writing.WriteError(w, errors.New("系统错误，请联系管理员!!"))
	}

	writing.WriteData(w, hr)
}
