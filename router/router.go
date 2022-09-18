package router

import (
	"blog/api"
	"blog/views"
	"net/http"
)

func Router() {
	//mvc 模式
	//页面
	http.HandleFunc("/", views.Html.Index)
	http.HandleFunc("/c/", views.Html.Category)
	http.HandleFunc("/login", views.Html.Login)
	http.HandleFunc("/p/", views.Html.Details)
	http.HandleFunc("/writing", views.Html.Writing)
	http.HandleFunc("/pigeonhole", views.Html.Pigeonhole)

	//接口
	http.HandleFunc("/api/v1/index", api.API.Viewer)
	http.HandleFunc("/api/v1/categorys", api.API.Categorys)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/api/v1/post/search", api.API.GetPostSearch)

	//静态资源
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

}
