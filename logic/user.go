package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// UserLoginLogic 用户登录逻辑业务处理
func UserLoginLogic(p *models.ParamLogin) (token string, err error) {
	user, err := mysql.UserLogin(p.Username, p.Password)
	if err != nil {
		return "", err
	}
	return jwt.GenToken(user.UserID, user.Username)
	/*token, err = jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", err
	}
	err = redis.SetToken(user.UserID, token)
	return token, err*/
}

// UserSignUpLogic 用户注册业务逻辑处理
func UserSignUpLogic(p *models.ParamSignUp) (err error) {
	//判断用户是否存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	//生成UID
	userID := snowflake.GenID()
	user := models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(&user)
}
