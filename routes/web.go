package routes

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Http"
	middle "gin-lmsail/app/Middleware"
	"gin-lmsail/app/Socket"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"html/template"
	"path/filepath"
)

func InitWebRouter() *gin.Engine {
	sessionKey := Helpers.ConfigOne("session.SESSION_SECRET", "gin-session")
	r := gin.Default()

	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"mbSubstr":     Helpers.MbSubstr,
		"formatAsDate": Helpers.FormatAsDate,
	})

	// 静态文件/HTML模板渲染
	r.Static("/statics", filepath.Join(Helpers.GetGoRunPath(), "resource"))
	r.StaticFile("/favicon.ico", filepath.Join(Helpers.GetGoRunPath(), "resource/favicon.ico"))
	r.LoadHTMLGlob(filepath.Join(Helpers.GetGoRunPath(), "views/**/*.html"))
	r.Use(EnableCookieSession(sessionKey), middle.SharedData())

	r.GET("/", Http.Index)
	r.GET("/about", Http.About)
	r.GET("/list", Http.Article)
	r.GET("/show/:id", Http.Show)
	r.GET("/time", Http.Time)
	r.GET("/login", Http.Login)
	r.GET("/logout", Http.Logout)
	r.GET("/register", Http.Register)
	r.POST("/login", Http.LoginAction)
	r.POST("/register", Http.RegisterAction)

	// 需要登录认证后才可访问的路由
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	Auth := r.Group("/", middle.UserAuth())
	{
		Auth.GET("/center/:uid", Http.Users)
		Auth.GET("/edit", Http.Edit)
		Auth.GET("/write", Http.Write)
		Auth.POST("/edit", Http.EditAction)
		Auth.POST("/write", Http.WritePost)
	}

	// websocket 在线瞎聊
	r.Any("/chat", Socket.ServerWs)

	// 接口类路由
	r.POST("/delblog", Http.DelBlogPost)

	return r
}

// Cookie 模式
func EnableCookieSession(SESSIONKEY string) gin.HandlerFunc {
	store := cookie.NewStore([]byte(SESSIONKEY))
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 86400, Path: "/"})
	return sessions.Sessions(SESSIONKEY, store)
}

// Redis 模式
func EnableRedisSession(SESSIONKEY string) gin.HandlerFunc {
	store, _ := redis.NewStore(10, "tcp", "redis:6379", "", []byte(SESSIONKEY))
	return sessions.Sessions(SESSIONKEY, store)
}
