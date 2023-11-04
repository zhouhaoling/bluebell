package controller

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

const CtxUserIDKey = "userID"

var ErrorUserNotLogin = errors.New("用户未登录")

// getCurrentUser 获取当前登录的用户ID
func getCurrentUserID(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(CtxUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return
}

// GetPageInfo 获取分页参数
func GetPageInfo(c *gin.Context) (int, int) {
	pageStr := c.Query("page")
	sizeStr := c.Query("size")
	var (
		err  error
		page int
		size int
	)
	page, err = strconv.Atoi(pageStr)
	if err != nil {
		page = 1
	}
	size, err = strconv.Atoi(sizeStr)
	if err != nil {
		size = 10
	}
	return page, size
}
