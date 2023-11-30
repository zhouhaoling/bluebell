package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"errors"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// UserLoginHandler 处理登录业务请求的函数
func UserLoginHandler(c *gin.Context) {
	var p models.ParamLogin
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("Login with invaild param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}
	//业务逻辑处理
	user, err := logic.UserLogin(&p)
	if err != nil {
		zap.L().Error("logic.UserLoginLogic failed", zap.String("username", p.UserName), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	//返回响应
	ResponseSuccess(c, gin.H{
		"user_id":       fmt.Sprintf("%d", user.UserID),
		"user_name":     user.UserName,
		"access_token":  user.AccessToken,
		"refresh_token": user.RefreshToken,
	})
}

// UserSignUpHandler 处理注册业务请求
func UserSignUpHandler(c *gin.Context) {
	var p models.ParamSignUp
	err := c.ShouldBindJSON(&p)
	if err != nil {
		zap.L().Error("SignUp with invaild param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, errs.Translate(trans))
		return
	}

	fmt.Println(p)
	if err = logic.UserSignUp(&p); err != nil {
		zap.L().Error("logic.UserSignUpLogic failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, nil)
}

// RefreshTokenHandler 刷新token
func RefreshTokenHandler(c *gin.Context) {
	//rt := c.Query("refresh_token")
	rt := c.GetHeader("refresh_token") //从请求头拿取refresh_token
	//fmt.Println("rt:", rt)
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		ResponseErrorWithMsg(c, CodeInvalidToken, "请求头缺少Auth Token")
		c.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		ResponseErrorWithMsg(c, CodeInvalidToken, "token格式不对")
		c.Abort()
		return
	}
	//fmt.Println("parts[1]:", parts[1])
	aToken, rToken, err := jwt.RefreshToken(parts[1], rt)
	if err != nil {
		zap.L().Error("jwt.RefreshToken() failed", zap.Error(err))
		c.Abort()
		return
	}
	ResponseSuccess(c, gin.H{
		"access_token":  aToken,
		"refresh_token": rToken,
	})
}
