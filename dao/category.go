package dao

import (
	"blog/models"
	"log"
)

func GetCategoryNameById(cId int) string {

	row := DB.QueryRow("select name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println("GetCategoryNameById 查询出错:", row.Err())
		return ""
	}

	var categoryName string
	_ = row.Scan(&categoryName)
	return categoryName
}

func GetAllCategory() ([]models.Category, error) {
	rows, err := DB.Query("select cid,name from blog_category")
	if err != nil {
		log.Println("GetAllCategory 查询出错:", err)
		return nil, err
	}
	var categorys []models.Category
	for rows.Next() {
		var category models.Category
		err = rows.Scan(&category.Cid, &category.Name)
		if err != nil {
			log.Println("GetAllCategory 取值出错", err)
			return nil, err
		}
		categorys = append(categorys, category)
	}
	return categorys, nil
}
