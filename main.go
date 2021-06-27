package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello，这里是go-blog!</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "about页面<a href='www.baidu.com'>www.baidu.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>页面未找到:(</h1>")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", router)
}
