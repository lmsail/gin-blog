package Service

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Models"
	"github.com/jinzhu/gorm"
)

func GetCategoryList() (list []*Models.Category, err error) {
	list, err = Models.GetCateList()
	if err != nil {
		Helpers.Fatal("查询分类列表异常：", err)
		return
	}
	return
}

func GetCateNameById(CateId int, defaultName string) string {
	cate := Models.Category{
		Model: gorm.Model{ID: uint(CateId)},
	}
	err := cate.GetCateInfo()
	if err != nil {
		Helpers.Fatal("查询分类列表异常：", err)
		return defaultName
	}
	return cate.CateName
}

// 分类信息
func GetCategoryInfo(CateId uint) (*Models.Category, error) {
	cate := &Models.Category{
		Model: gorm.Model{ID: uint(CateId)},
	}
	err := cate.GetCateInfo()
	if err != nil {
		return cate, err
	}
	return cate, nil
}
