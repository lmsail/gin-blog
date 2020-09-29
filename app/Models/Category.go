package Models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	CateName string `gorm:"index:cate_name"`
}

// 获取分类列表
func GetCateList() (list []*Category, err error) {
	err = DB.Order("id Desc").Find(&list).Error
	return list, err
}

// 获取指定分类信息
func (cate *Category) GetCateInfo() error {
	return DB.Where("id = ?", cate.ID).First(&cate).Error
}
