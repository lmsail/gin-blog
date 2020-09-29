package Models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title   string `gorm:"index:title"`
	UserId  uint   `gorm:"index:user_id"`
	Content string
	CateId  uint
	View    int   `gorm:"default:'0'"`
	Comment int   `gorm:"default:'0'"`
	IsTop   uint8 `gorm:"default:'0'"`
}

// 文章、对应文章分类信息
type ArticleRecord struct {
	Article
	Category
}

// 展示文章详情数据
type ArticleShow struct {
	Article
	Category
	UserInfo
}

// 发布文章
func (article *Article) CreateBlog() error {
	return DB.FirstOrCreate(&article, "title = ? or content = ?", article.Title, article.Content).Error
}

// 获取指定文章信息
func (article *Article) ShowArticleInfo() error {
	return DB.Where("id = ?", article.ID).Find(&article).Error
}

// 查询置顶的文章列表
func GetTopArticleList(nums int) (list []*Article, err error) {
	//err = DB.Where("is_top = ?", 1).Limit(nums).Find(&list).Error
	err = DB.Limit(nums).Find(&list).Error
	return list, err
}

// 查询博客列表 分页
func GetArticleList(cateId, UserId, page int) ([]*Article, int, error) {
	db := DB
	var (
		list      []*Article
		totalNums int
	)
	if cateId > 0 {
		db = db.Where("cate_id = ?", cateId)
	}
	if UserId > 0 {
		db = db.Where("user_id = ?", UserId)
	}
	db.Find(&list).Count(&totalNums)
	if page > 0 {
		db = db.Limit(PageSize).Offset((page - 1) * PageSize)
	} else {
		db = db.Limit(PageSize).Offset(0)
	}
	err := db.Order("id DESC").Find(&list).Error
	return list, totalNums, err
}

// 文章浏览量+1
func (article *Article) AddArticleView(id int) error {
	return DB.Model(&article).Update("view", gorm.Expr("view + ?", 1)).Error
}

// 删除文章
func (article *Article) DeleteArticle() error {
	return DB.Delete(&article).Error
}
