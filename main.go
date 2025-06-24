package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// 定义一个全局的 map 存储所有路径和对应响应
var mockData map[string]interface{}

// 读取 JSON 文件中的内容到 map 中
func loadMockData(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(&mockData)
}

// 请求处理函数
func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	if data, ok := mockData[path]; ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	} else {
		http.NotFound(w, r)
	}
}

// 启动主函数
func main() {
	err := loadMockData("mock_data.json")
	if err != nil {
		log.Fatalf("加载 mock_data.json 失败: %v\n", err)
	}

	port := 4332
	fmt.Printf("✅ Mock 服务启动成功，监听端口 %d\n", port)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
