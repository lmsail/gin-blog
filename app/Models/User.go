package Models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Username        string `gorm:"type:varchar(255);not null"`
	Nickname        string `gorm:"type:varchar(255)"`
	Password        string `gorm:"type:varchar(64);not null"`
	Avatar          string `gorm:"type:varchar(255)"`
	Lockstate       int8   `gorm:"default:'0'"`
	Autograph       string `gorm:"autograph:varchar(255)"`
	Introduction    string `gorm:"introduction:varchar(255)"`
	PersonalWebsite string `gorm:"personal_website:varchar(255)"`
}

// 用户基础信息，用于数据关联
type UserInfo struct {
	ID        int
	Username  string
	Nickname  string
	Lockstate int8
	Avatar    string
}

// 用户注册
func (user *Users) Register() error {
	return DB.FirstOrCreate(&user, "username = ?", user.Username).Error
}

// 登录验证
func (user *Users) Signin() error {
	return DB.First(&user, "username = ? AND password = ?", user.Username, user.Password).Error
}

// 获取单个用户信息
func (user *Users) GetUserInfo() error {
	return DB.First(&user, "id = ?", user.ID).Error
}

// 获取用户基础信息
func (user *UserInfo) GetUserBaseInfo() error {
	return DB.Table("users").Select("id, username, nickname, avatar, lockstate").Where("id = ?", user.ID).Scan(&user).Error
}

// 更新用户信息
func (user *Users) SaveUserInfo(userInfo map[string]string) error {
	return DB.Model(user).Updates(userInfo).Error
}
