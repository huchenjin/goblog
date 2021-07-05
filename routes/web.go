package routes

import (
	"goblog/app/http/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterWebRoutes 注册网页相关路由
func RegisterWebRoutes(r *mux.Router) {
	pagesController := new(controllers.PagesController)
	// 静态页面
	r.HandleFunc("/", pagesController.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pagesController.About).Methods("GET").Name("about")
	r.NotFoundHandler = http.HandlerFunc(pagesController.NotFound)
}
