package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"encoding/hex"
	"errors"

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
func UserLogin(username string, password string) (user *models.User, err error) {
	result := db.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, ErrorUserNotExist
	}
	//查询数据库失败
	if err != nil {
		return nil, err
	}
	if user.Password != encryptPassword(password) {
		return nil, ErrorInvalidPassword
	}
	return user, nil
}

// encryptPassword 密码加密
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// GetUserById 根据用户id查询数据
func GetUserById(userID int64) (*models.User, error) {
	result := db.Where("user_id = ?", userID).First(&user)
	return &user, result.Error
}
