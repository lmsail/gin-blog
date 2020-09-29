package Service

import (
	"errors"
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Models"
	"github.com/jinzhu/gorm"
)

// 组装返回数据，有点冗余，待功力提升回头优化
// 循环组装文章以及所属分类可使用sql 联查优化
type ReturnData struct {
	CateList       []*Models.Category
	ArticleTopList []*Models.Article
	ArticleList    []*Models.ArticleRecord
	Paginator      map[string]interface{}
}

// 组合文章列表、分类列表带分页
func HandleArticleList(CateId, UserId, Page int) (data *ReturnData) {
	var list []*Models.ArticleRecord
	cateList, err := GetCategoryList()
	topList, err := Models.GetTopArticleList(10)
	if err != nil {
		Helpers.Fatal("获取置顶文章列表异常：", err)
		return
	}
	articleList, total, err := Models.GetArticleList(CateId, UserId, Page)
	if err != nil {
		Helpers.Fatal("获取文章列表异常：", err)
		return
	}
	// 组合文章与分类信息
	for _, article := range articleList {
		// 将文章信息写入
		recordInfo := &Models.ArticleRecord{
			Article: *article,
		}
		// 将对应的分类信息写入
		for _, category := range cateList {
			if article.CateId == category.ID {
				recordInfo.Category = *category
				break
			}
		}
		list = append(list, recordInfo)
	}
	paginator := Helpers.Paginator(Page, Models.PageSize, total)
	return &ReturnData{
		CateList:       cateList,
		ArticleTopList: topList,
		ArticleList:    list,
		Paginator:      paginator,
	}
}

// 查询单个文章信息
func GetArticleDetail(id int) (details *Models.ArticleShow, err error) {
	article := &Models.Article{
		Model: gorm.Model{ID: uint(id)},
	}
	err = article.ShowArticleInfo()
	if err != nil {
		Helpers.Fatal("查询文章信息出错，文章ID："+string(id), err)
		return
	} // 文章信息
	cateInfo, err := GetCategoryInfo(article.CateId)
	if err != nil {
		return
	} // 分类信息
	uinfo, err := GetUserBaseInfo(int(article.UserId))
	if err != nil {
		return
	} // 用户信息
	err = article.AddArticleView(id)
	if err != nil {
		return
	} // 浏览量+1
	return &Models.ArticleShow{
		Article:  *article,
		Category: *cateInfo,
		UserInfo: *uinfo,
	}, nil // 组装返回信息
}

func CreateArticle(Title, Content string, UserId, CateId uint) error {
	article := Models.Article{
		Title:   Title,
		UserId:  UserId,
		Content: Content,
		CateId:  CateId,
	}
	if _, err := GetCategoryInfo(CateId); err != nil {
		return errors.New("分类不存在")
	}
	if err := article.CreateBlog(); err != nil {
		return err
	}
	return nil
}

func CheckArticleAuthor(ArticleID, UserId uint) bool {
	article := &Models.Article{
		Model: gorm.Model{ID: ArticleID},
	}
	err := article.ShowArticleInfo()
	if err != nil {
		return false
	}
	if article.UserId != UserId {
		return false
	}
	return true
}

func DeleteArticleById(ArticleID uint) bool {
	article := &Models.Article{
		Model: gorm.Model{ID: ArticleID},
	}
	err := article.DeleteArticle()
	if err != nil {
		return false
	}
	return true
}
