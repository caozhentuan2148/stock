package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caozhentuan2148/stock/internal/handler"
)

func main() {
	fmt.Println("股票分析与资产配置模拟软件启动中...")

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/health", handler.HealthHandler)

	port := ":8080"
	log.Printf("服务器启动在端口 %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}
