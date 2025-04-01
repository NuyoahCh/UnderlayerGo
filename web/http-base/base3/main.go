// Package base3
// @Author NuyoahCh
// @Date 2025/2/12 23:13
// @Desc
package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Headler[%q] = %q\n", k, v)
		}
	})
	r.Run(":9999")
}
