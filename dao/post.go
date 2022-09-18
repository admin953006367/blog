package dao

import (
	"blog/models"
	"log"
)

func CountGetAllPost() (count int) {
	rows := DB.QueryRow("select count(1) from blog_post")
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostBySlug(slug string) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where slug = ?", slug)
	_ = rows.Scan(&count)
	return
}

func CountGetAllPostByCategoryId(cid int) (count int) {
	rows := DB.QueryRow("select count(1) from blog_post where category_id = ?", cid)
	_ = rows.Scan(&count)
	return
}

func GetPostPage(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select  * from blog_post Order by  create_at desc limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetPostPageBySlug(slug string, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select  * from blog_post where slug = ? Order by create_at desc limit ?,?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetCategoryPostPage(pid, page, limit int) ([]models.Post, error) {
	page = (page - 1) * limit
	rows, err := DB.Query("select  * from blog_post where category_id = ? limit ?,? ", pid, page, limit)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post

		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, err
}

func GetPostById(pid int) (*models.Post, error) {
	row := DB.QueryRow("select * from blog_post where pid = ?", pid)
	if row.Err() != nil {
		return nil, row.Err()
	}

	var post = &models.Post{}
	err := row.Scan(&post.Pid, &post.Title, &post.Content, &post.Markdown, &post.CategoryId, &post.UserId, &post.ViewCount, &post.Type, &post.Slug, &post.CreateAt, &post.UpdateAt)
	if err != nil {
		return nil, err
	}
	return post, err
}

func SavePost(post *models.Post) {
	ret, err := DB.Exec("insert into blog_post "+
		"(title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) "+
		"values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt,
	)
	if err != nil {
		log.Println(err)
	}
	pid, _ := ret.LastInsertId()
	post.Pid = int(pid)
}

func UpdatePost(post *models.Post) {
	_, err := DB.Exec("update blog_post set title=?,content=?,markdown=?,category_id=?,type=?,slug=?,update_at=? where pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.Pid,
	)
	if err != nil {
		log.Println(err)
	}
}

func GetPostAll() ([]models.Post, error) {
	rows, err := DB.Query("select  * from blog_post")
	if err != nil {
		return nil, err
	}

	var posts = []models.Post{}
	for rows.Next() {
		var post = models.Post{}
		err = rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, err
}

func GetPostBySearch(condition string) ([]models.SearchResp, error) {
	rows, err := DB.Query("select pid,title from blog_post where title like ? Order by create_at desc limit 5 ", "%"+condition+"%")
	if err != nil {
		return nil, err
	}

	var posts = []models.SearchResp{}
	for rows.Next() {
		var post = models.SearchResp{}
		err = rows.Scan(
			&post.Pid,
			&post.Title,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, err
}
