package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
)

func SavePost(post *models.Post) {
	dao.SavePost(post)
}

func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}

func GetPostDetails(pid int) (*models.PostRes, error) {

	post, err := dao.GetPostById(pid)

	if err != nil {
		return nil, err
	}

	categoryName := dao.GetCategoryNameById(post.CategoryId)
	userName := dao.GetUserNameById(post.UserId)

	postMore := models.PostMore{
		Pid:          post.Pid,
		Title:        post.Title,
		Slug:         post.Slug,
		Content:      template.HTML(post.Content),
		CategoryId:   post.CategoryId,
		CategoryName: categoryName,
		UserId:       post.UserId,
		UserName:     userName,
		ViewCount:    post.ViewCount,
		Type:         post.Type,
		CreateAt:     models.DateDay(post.CreateAt),
		UpdateAt:     models.DateDay(post.UpdateAt),
	}

	var postRes = &models.PostRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Article:      postMore,
	}
	return postRes, nil

}

func GetWriting() (*models.WritingRes, error) {

	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	var writingRes = &models.WritingRes{
		Title:     config.Cfg.Viewer.Title,
		Categorys: categorys,
		CdnURL:    config.Cfg.System.CdnURL,
	}
	return writingRes, nil
}
