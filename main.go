package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello，这里是go-blog!</h1>")
	} else if r.URL.Path == "/about" {
		fmt.Fprint(w, "about页面")
	} else {
		fmt.Fprint(w, "<h1>页面未找到:(</h1>")
	}
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe(":3000", nil)
}
