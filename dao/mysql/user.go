package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var user models.User

const secret = "lonlyyahh.com"

// InsertUser 向数据库中创建一条用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = encryptPassword(user.Password)
	result := db.Create(&user)
	if result.RowsAffected > 0 {
		return nil
	}
	return result.Error
}

// CheckUserExist 查找用户是否存在
func CheckUserExist(username string) (err error) {
	result := db.Where("username = ?", username).First(&user)
	//排除没有找到记录的错误
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return err
	}
	if result.RowsAffected > 0 {
		return ErrorUserExist
	}
	return
}

// UserLogin 用户登录
func UserLogin(user *models.User) (err error) {
	LPassword := user.Password //记录登录的密码
	user.Password = "idns"
	result := db.Where("username = ?", user.UserName).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return ErrorUserNotExist
	}
	//判断密码
	fmt.Println(user.Password)
	if user.Password != encryptPassword(LPassword) {
		return ErrorInvalidPassword
	}
	return nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// GetUserById 根据用户id查询数据
func GetUserById(userID int64) (user *models.User, err error) {
	result := db.Where("user_id = ?", userID).First(&user)
	return user, result.Error
}
