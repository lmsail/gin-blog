package Http

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LoginPost struct {
	UserName string `form:"username" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

type RegisterPost struct {
	UserName string `form:"username" binding:"required"`
	NickName string `form:"nickname" binding:"required"`
	PassWord string `form:"password" binding:"required"`
}

type EditPost struct {
	Nickname        string `form:"nickname" binding:"required"`
	Autograph       string `form:"autograph" binding:"required"`
	Introduction    string `form:"introduction" binding:"required"`
	PersonalWebsite string `form:"personal_website" binding:"required"`
}

// 用户中心
func Users(c *gin.Context) {
	assign := AssignData(c, "会员中心")
	userId, _ := strconv.Atoi(c.Param("uid"))
	page, _ := strconv.Atoi(c.Query("page"))
	if userId <= 0 {
		c.HTML(http.StatusBadRequest, "error/400", nil)
		return
	}
	// 查询用户信息
	user, err := Service.GetUserInfo(userId)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500", nil)
		return
	}
	data := Service.HandleArticleList(0, userId, page)
	assign["paginator"] = data.Paginator
	assign["list"] = data.ArticleList
	assign["user"] = user
	c.HTML(http.StatusOK, "users/index", assign)
}

func Edit(c *gin.Context) {
	assign := AssignData(c, "编辑个人资料")
	c.HTML(http.StatusOK, "users/edit", assign)
}

// 用户登录
func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login/login", AssignData(c, "会员登录"))
}

// 用户注册
func Register(c *gin.Context) {
	c.HTML(http.StatusOK, "login/register", AssignData(c, "会员注册"))
}

// 退出登录
func Logout(c *gin.Context) {
	Helpers.ClearSession(c)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 用户资料修改
func EditAction(c *gin.Context) {
	var edit EditPost
	assign := AssignData(c, "编辑个人资料")
	if err := c.ShouldBind(&edit); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "users/edit", assign)
	}
	file, _ := c.FormFile("avatar")
	userId, exists := c.Get("UserId")
	if !exists {
		c.Redirect(http.StatusMovedPermanently, "/login")
		return
	}
	err := Service.SaveUserInfo(c, userId.(uint), gin.H{
		"avatar":           file,
		"nickname":         edit.Nickname,
		"autograph":        edit.Autograph,
		"introduction":     edit.Introduction,
		"personal_website": edit.PersonalWebsite,
	})
	if err != "" {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "users/edit", assign)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/edit")
	}
}

// 用户登录动作
func LoginAction(c *gin.Context) {
	var login LoginPost
	assign := AssignData(c, "会员登录")
	if err := c.ShouldBind(&login); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/login", assign)
		return
	}
	UserId, err := Service.Login(login.UserName, Helpers.Md5(login.PassWord))
	if err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/login", assign)
		return
	}
	Helpers.SetSession(c, "UserId", UserId)
	c.Redirect(http.StatusMovedPermanently, "/")
}

// 用户注册动作
func RegisterAction(c *gin.Context) {
	var register RegisterPost
	assign := AssignData(c, "会员注册")
	if err := c.ShouldBind(&register); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "login/register", assign)
		return
	}
	err := Service.Register(register.UserName, Helpers.Md5(register.PassWord), register.NickName)
	if err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusOK, "login/register", assign)
	} else {
		c.Redirect(http.StatusMovedPermanently, "/login")
	}
}
