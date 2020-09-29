package Service

import (
	"gin-lmsail/app/Helpers"
	"gin-lmsail/app/Models"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"path/filepath"
	"strings"
	"errors"
)

func Login(UserName, PassWord string) (uint, error) {

	user := Models.Users{
		Username: UserName, Password: PassWord,
	}
	if err := user.Signin(); err != nil {
		return 0, err
	}
	if user.Lockstate == 1 {
		return 0, errors.New("该账户已被限制登录")
	}
	return user.ID, nil
}

func Register(UserName, PassWord, Nickname string) error {

	user := Models.Users{
		Username: UserName, Nickname: Nickname, Password: PassWord,
	}
	if err := user.Register(); err != nil {
		return err
	}
	return nil
}

func GetUserBaseInfo(user_id int) (user *Models.UserInfo, err error) {
	user = &Models.UserInfo{ID: user_id}
	err = user.GetUserBaseInfo()
	if err != nil {
		return
	}
	return
}

func GetUserInfo(user_id int) (user Models.Users, err error) {
	user.ID = uint(user_id)
	err = user.GetUserInfo()
	return user, err
}

func UploadUserAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	savePath := ""
	if file != nil {
		savePath = filepath.Join(Helpers.GetGoRunPath(), "resource/imgs/avatar/", Helpers.GetRandomString(10)+".png")
		err := c.SaveUploadedFile(file, savePath)
		if err != nil {
			return "", err
		}
		savePath = "/statics/imgs" + strings.Split(savePath, "imgs")[1]
	}
	return savePath, nil
}

func SaveUserInfo(c *gin.Context, userId uint, data gin.H) string {
	savePath, err := UploadUserAvatar(c, data["avatar"].(*multipart.FileHeader))
	if err != nil {
		return "头像上传失败"
	}
	user := &Models.Users{}
	user.ID = userId
	saveData := map[string]string{
		"nickname":         data["nickname"].(string),
		"avatar":           savePath,
		"autograph":        data["autograph"].(string),
		"introduction":     data["introduction"].(string),
		"personal_website": data["personal_website"].(string),
	}
	if savePath == "" {
		delete(saveData, "avatar")
	}
	err = user.SaveUserInfo(saveData)
	if err != nil {
		Helpers.Fatal("更新用户信息失败", err)
		return "用户资料修改失败"
	}
	return ""
}
