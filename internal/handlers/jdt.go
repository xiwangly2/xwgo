package handlers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var (
	mu       sync.Mutex
	progress int
)

func JdtHandler(w http.ResponseWriter, r *http.Request) {
	// 设置CORS标头以允许跨域请求
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // 允许的HTTP方法
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type") // 允许的HTTP标头

	if r.Method == "OPTIONS" {
		// 处理预检请求（OPTIONS请求）
		w.WriteHeader(http.StatusOK)
		return
	}

	// 获取当前时间的秒数
	_ = time.Now().Second()

	// 生成一个0到5之间的随机递增步长
	increment := rand.Intn(6)

	mu.Lock()
	defer mu.Unlock()

	// 增加进度，并确保不超过100
	progress += increment
	if progress > 100 {
		progress = 100
	}

	// 创建一个包含进度值的 JSON 对象
	progressData := map[string]interface{}{
		"code":     "200",
		"progress": progress,
	}

	// 将 JSON 对象编码为 JSON 格式
	dataJSON, err := json.Marshal(progressData)
	if err != nil {
		http.Error(w, "JSON encoding error", http.StatusInternalServerError)
		return
	}

	// 设置响应头部
	w.Header().Set("Content-Type", "application/json")

	// 返回 JSON 数据
	_, err = w.Write(dataJSON)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
