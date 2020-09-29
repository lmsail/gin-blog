package Middleware

import (
	"gin-lmsail/app/Helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 前台用户认证中间件
func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userId := Helpers.GetSession(c, "UserId"); userId != nil {
			c.Next()
			return
		}
		c.HTML(http.StatusBadRequest, "login/login", gin.H{
			"errorMsg": "请先登录!",
		})
		c.Abort()
	}
}

// 接口类路由验证用户是否登录
func UserApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userId := Helpers.GetSession(c, "UserId"); userId != nil {
			c.Next()
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"code": http.StatusBadRequest,
			"msg":  "用户未登录",
			"data": nil,
		})
		c.Abort()
	}
}

// 注入用户信息
func SharedData() gin.HandlerFunc {
	return func(c *gin.Context) {
		if userId := Helpers.GetSession(c, "UserId"); userId != nil {
			c.Set("UserId", userId)
		}
		c.Next()
	}
}
