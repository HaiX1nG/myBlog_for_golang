package router

import (
	"awesomeProject/src/models"
	"awesomeProject/src/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	// 使用cors中间件
	r.Use(utils.Cors())

	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"text": "hello world!",
		})
	})

	r.GET("/user", models.GetUserInfo)

	r.POST("/login", models.LoginHandle)

	r.POST("/register", models.RegisterHandle)

	return r
}
