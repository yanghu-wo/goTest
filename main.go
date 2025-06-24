package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

// mock 数据结构：map[path]map[key]内容
var mockData map[string]map[string]interface{}

func loadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	err = json.NewDecoder(file).Decode(&config)
	return &config, err
}

func loadMockData(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&mockData)
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	query := r.URL.Query()
	systemCode := query.Get("systemCode")

	dataForPath, ok := mockData[path]
	if !ok {
		http.NotFound(w, r)
		return
	}

	resp, ok := dataForPath[systemCode]
	if !ok {
		resp = dataForPath["default"]
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// 加载 config.json
	config, err := loadConfig("config.json")
	if err != nil {
		log.Fatalf("加载配置文件失败: %v\n", err)
	}

	// 加载 mock_data.json
	err = loadMockData("mock_data.json")
	if err != nil {
		log.Fatalf("加载mock数据失败: %v\n", err)
	}

	addr := fmt.Sprintf("%s:%d", config.Host, config.Port)
	fmt.Printf("✅ MockServer 启动，监听地址 %s\n", addr)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
