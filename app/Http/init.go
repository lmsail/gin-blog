package Http

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Blog map[string]string

func init() {
	Blog = Helpers.ConfigMultiple("blog")
}

// 绑定全局返回值
func AssignData(c *gin.Context, title string) gin.H {
	var user Models.Users
	if UserId, ok := c.Get("UserId"); ok && UserId != nil {
		user.ID = UserId.(uint)
		err := user.GetUserInfo()
		if err != nil {
			c.Redirect(http.StatusMovedPermanently, "error/500")
		}
	}
	return gin.H{
		"self":     user,
		"app_name": Blog["APP_NAME"],
		"title":    title + " - " + Blog["APP_NAME"],
	}
}
