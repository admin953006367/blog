package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
	"html/template"
)

func GetCategoryIndexInfo(cid, page, limit int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err2 := dao.GetCategoryPostPage(cid, page, limit)
	if err2 != nil {
		return nil, err2
	}

	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)

		//转为中文字符串
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}

		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}

		postMores = append(postMores, postMore)
	}

	total := dao.CountGetAllPostByCategoryId(cid)

	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}

	categorysName := dao.GetCategoryNameById(cid)

	//数据内容必须定义
	var home = &models.HomeResponse{
		Viewer:    config.Cfg.Viewer,
		Categorys: categorys,
		Posts:     postMores,
		Total:     total,
		Page:      page,
		Pages:     pages,
		PageEnd:   page != pagesCount,
	}

	var hr = &models.CategoryResponse{
		home,
		categorysName,
	}

	return hr, nil

}
