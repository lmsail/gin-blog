package Http

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticlePost struct {
	Title   string `form:"title" binding:"required"`
	UserId  int    `form:"user_id"`
	CateId  uint   `form:"cate_id" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func Article(c *gin.Context) {
	assign := AssignData(c, "博客列表")
	cateId, _ := strconv.Atoi(c.Query("cateId"))
	page, _ := strconv.Atoi(c.Query("page"))
	data := Service.HandleArticleList(cateId, 0, page)
	assign["paginator"] = data.Paginator
	assign["cateList"] = data.CateList
	assign["topList"] = data.ArticleTopList
	assign["list"] = data.ArticleList
	assign["friendLinkList"] = gin.H{
		"LmSail博客":     "http://www.lmsail.com",
		"Gin-LmSail博客": "http://go.lmsail.com",
	}
	assign["cateName"] = "博客列表"
	if cateId > 0 {
		assign["cateName"] = Service.GetCateNameById(cateId, "博客列表")
	}
	c.HTML(http.StatusOK, "article/list", assign)
}

func Show(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	details, err := Service.GetArticleDetail(id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500", nil)
	} else {
		assign := AssignData(c, details.Title)
		assign["details"] = details
		c.HTML(http.StatusOK, "article/show", assign)
	}
}

func Write(c *gin.Context) {
	assign := AssignData(c, "写文章")
	list, err := Service.GetCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500", nil)
	} else {
		assign["list"] = list
		c.HTML(http.StatusOK, "article/write", assign)
	}
}

func WritePost(c *gin.Context) {
	var formParams ArticlePost
	assign := AssignData(c, "写文章")
	if err := c.ShouldBind(&formParams); err != nil {
		assign["errorMsg"] = err
		c.HTML(http.StatusBadRequest, "article/write", assign)
	} else {
		UserId, ok := c.Get("UserId")
		if !ok {
			assign["errorMsg"] = "检测您未登录，请先登录"
			c.HTML(http.StatusBadRequest, "article/write", assign)
		}
		err := Service.CreateArticle(formParams.Title, formParams.Content, UserId.(uint), formParams.CateId)
		if err != nil {
			assign["errorMsg"] = err
			c.HTML(http.StatusBadRequest, "article/write", assign)
		} else {
			c.Redirect(http.StatusMovedPermanently, "/list")
		}
	}
}

func DelBlogPost(c *gin.Context) {
	articleID, _ := strconv.Atoi(c.PostForm("id"))
	UserId, exists := c.Get("UserId")
	if !exists {
		c.JSON(http.StatusBadRequest, Helpers.Json(http.StatusBadRequest, "用户未登录", nil))
		return
	}
	if status := Service.CheckArticleAuthor(uint(articleID), UserId.(uint)); status {
		if delStatus := Service.DeleteArticleById(uint(articleID)); delStatus {
			c.JSON(http.StatusOK, Helpers.Json(http.StatusOK, "已删除", nil))
			return
		}
		c.JSON(http.StatusBadRequest, Helpers.Json(http.StatusBadRequest, "删除失败", nil))
	} else {
		c.JSON(http.StatusBadRequest, Helpers.Json(http.StatusBadRequest, "权限不足，越权删除", nil))
	}
}
