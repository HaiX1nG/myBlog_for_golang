package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 跨域请求设置

var CorsConfig = cors.Config{
	// 允许哪里传来的请求，生产环境可以直接替换为域名
	AllowOrigins: []string{"http://localhost:5173"},
	// 允许的HTTP方法列表，如 GET、POST、PUT等。默认为["*"]（全部允许）
	AllowMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
	// 允许的HTTP头部列表。默认为["*"]（全部允许）
	AllowHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept", "token", "x-custom-header"},
	// 是否允许浏览器发送Cookie。默认为false
	AllowCredentials: true,
	// 预检请求（OPTIONS）的缓存时间（秒）。默认为5分钟
	MaxAge: 60 * 5,
}

// 跨域请求接口

func Cors() gin.HandlerFunc {
	return cors.New(CorsConfig)
}
