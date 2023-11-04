package controller

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/jwt"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// LoginHandler 处理登录业务请求的函数
func LoginHandler(c *gin.Context) {
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
	token, err := logic.UserLoginLogic(&p)
	if err != nil {
		zap.L().Error("logic.UserLoginLogic failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		if errors.Is(err, jwt.ErrorInvalidToken) {
			ResponseError(c, CodeServerBusy)
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	ResponseSuccess(c, token)
}

// SignUpHandler 处理注册业务请求
func SignUpHandler(c *gin.Context) {
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
	//手动对请求参数进行详细校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	//	zap.L().Error("SignUp with invaild param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	fmt.Println(p)
	if err = logic.UserSignUpLogic(&p); err != nil {
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
