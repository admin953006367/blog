package main

import (
	"blog/common"
	"blog/router"
	"fmt"
	"net/http"
)

func init() {
	//模板加载
	common.LoadTemplate()
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}

	//路由
	router.Router()

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
