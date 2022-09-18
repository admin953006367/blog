package service

import (
	"blog/config"
	"blog/dao"
	"blog/models"
)

func FindPostPigeonhole() (*models.PigeonholeRes, error) {

	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err2 := dao.GetPostAll()
	if err2 != nil {
		return nil, err2
	}

	pigeonholeMap := make(map[string][]models.Post)

	for _, post := range posts {
		at := post.CreateAt
		month := at.Format("2006-01")
		pigeonholeMap[month] = append(pigeonholeMap[month], post)
	}

	PigeonholeRes := &models.PigeonholeRes{
		Viewer:       config.Cfg.Viewer,
		SystemConfig: config.Cfg.System,
		Categorys:    categorys,
		Lines:        pigeonholeMap,
	}

	return PigeonholeRes, err2
}
