package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"net/http"
)

var router *mux.Router
var db *sql.DB

// Article  对应一条文章数据
type Article struct {
	Title, Body string
	ID          int64
}

func main() {
	// 初始化mysql
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()

	router = bootstrap.SetupRoute()

	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	router.Use(middlewares.ForceHTML)
	http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
}
