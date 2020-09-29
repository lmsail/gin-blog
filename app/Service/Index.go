package Service

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Models"
)

func GetIndexShow() (list []*Models.Article, err error) {
	list, err = Models.GetTopArticleList(6)
	if err != nil {
		Helpers.Fatal("获取置顶文章列表失败：", err)
	}
	return
}
