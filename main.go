package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response map[string]interface{}

var routes = map[string]func(http.ResponseWriter, *http.Request){
	"/test": func(w http.ResponseWriter, r *http.Request) {
		// 获取所有GET参数
		params := r.URL.Query()

		resp := Response{
			"message": "请求成功了",
			"params":  params,
		}

		writeJSON(w, http.StatusOK, resp)
	},
	"/login": func(w http.ResponseWriter, r *http.Request) {
		resp := Response{
			"code":    200,
			"message": "登录成功",
		}
		writeJSON(w, http.StatusOK, resp)
	},
}

func writeJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if h, ok := routes[r.URL.Path]; ok {
		h(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func main() {
	port := 4332
	fmt.Printf("MockServer启动，监听端口 %d\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
