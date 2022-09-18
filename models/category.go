package models

// 读取数据库中的内容

type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}
