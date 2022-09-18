package api

import (
	"blog/common"
	"blog/config"
	"blog/dao"
	"blog/models"
	"blog/service"
	"blog/utils"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

// Viewer 获取基础信息
func (*Api) Viewer(w http.ResponseWriter, r *http.Request) {
	viewr := config.Cfg.Viewer
	jsonStr, _ := json.Marshal(viewr)
	_, err := w.Write(jsonStr)
	if err != nil {
		log.Panicln("Categorys出错：", err)
	}
}

// Categorys 获取分类信息
func (*Api) Categorys(w http.ResponseWriter, r *http.Request) {
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	jsonStr, _ := json.Marshal(categorys)
	_, err := w.Write(jsonStr)
	if err != nil {
		log.Panicln("Categorys出错：", err)
	}
}

func (*Api) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	pIdStr := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(pIdStr)
	if err != nil {
		common.Error(w, errors.New("不识别此请求路径"))
		return
	}
	post, err := dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, post)
}

func (*Api) SaveAndUpdatePost(w http.ResponseWriter, r *http.Request) {

	//获取用户ID
	token := r.Header.Get("Authorization")
	_, claim, err := utils.ParseToken(token)
	if err != nil {
		common.Error(w, errors.New("登录已过期"))
	}
	uid := claim.Uid
	//post save
	method := r.Method
	switch method {
	case http.MethodPost:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(string)
		categoryId, _ := strconv.Atoi(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := float64(0)
		if params["type"] != nil {
			postType = params["type"].(float64)
		}

		pType := int(postType)

		post := &models.Post{
			Pid:        -1,
			Title:      title,
			Slug:       slug,
			Content:    content,
			Markdown:   markdown,
			CategoryId: categoryId,
			UserId:     uid,
			Type:       pType,
			CreateAt:   time.Now(),
			UpdateAt:   time.Now(),
		}
		service.SavePost(post)
		common.Success(w, post)
	case http.MethodPut:
		params := common.GetRequestJsonParam(r)
		cId := params["categoryId"].(float64)
		categoryId := int(cId)
		content := params["content"].(string)
		markdown := params["markdown"].(string)
		slug := params["slug"].(string)
		title := params["title"].(string)
		postType := float64(0)
		if params["type"] != nil {
			postType = params["type"].(float64)
		}
		pType := int(postType)
		pid := int(params["pid"].(float64))

		post := &models.Post{
			pid,
			title,
			slug,
			content,
			markdown,
			categoryId,
			uid,
			0,
			pType,
			time.Now(),
			time.Now(),
		}
		service.UpdatePost(post)
		common.Success(w, post)
	}

}

func (*Api) GetPostSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	condition := r.Form.Get("val")
	posts, _ := dao.GetPostBySearch(condition)
	var searchResps []models.SearchResp
	for _, post := range posts {
		searchResps = append(searchResps, models.SearchResp{
			post.Pid,
			post.Title,
		})
	}
	common.Success(w, searchResps)
}
