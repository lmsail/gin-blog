package Helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 写入session
func SetSession(c *gin.Context, key string, value interface{}) {
	session := sessions.Default(c)
	session.Set(key, value)
	_ = session.Save()
}

// 获取指定session key
func GetSession(c *gin.Context, key string) interface{} {
	session := sessions.Default(c)
	return session.Get(key)
}

// 清除session
func ClearSession(c *gin.Context) {
	s := sessions.Default(c)
	c.Set("UserId", nil)
	s.Clear()
	_ = s.Save()
}
