package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// UserLogin 用户登录逻辑业务处理
func UserLogin(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		UserName: p.UserName,
		Password: p.Password,
	}
	err = mysql.UserLogin(user)
	if err != nil {
		return nil, err
	}

	//return jwt.GenToken(user.UserID, user.Username)
	accessToken, refreshToken, err := jwt.GenToken(user.UserID, user.UserName)
	if err != nil {
		return
	}
	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	return
}

// UserSignUp 用户注册业务逻辑处理
func UserSignUp(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	if err := mysql.CheckUserExist(p.UserName); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	user := models.User{
		UserID:   userID,
		UserName: p.UserName,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(&user)
}
