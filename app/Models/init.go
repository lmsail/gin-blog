package Models

import (
	"gin-lmsail/app/Helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB       *gorm.DB
	PageSize int
)

func init() {
	var err error
	conf := Helpers.ConfigMultiple("mysql")
	DB, err = gorm.Open(conf["DB_CONNECTION"], conf["DB_USERNAME"]+":"+conf["DB_PASSWORD"]+"@tcp("+conf["DB_HOST"]+":"+conf["DB_PORT"]+")/"+conf["DB_DATABASE"]+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		Helpers.Fatal("数据库连接异常")
		panic("数据库连接异常")
	}
	PageSize = 5
	DB.AutoMigrate(&Users{}, &Article{}, &Category{})
}
