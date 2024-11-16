package middlewares

import (
	"net/http"
)

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 设置 CORS 头
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")                  // 允许的来源
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT,PATCH, DELETE, OPTIONS") // 允许的方法
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")           // 允许的请求头

		// 处理预检请求
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
