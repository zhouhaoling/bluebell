package controller

import (
	"bluebell/dao/redis"
	"bluebell/logic"
	"bluebell/models"

	"go.uber.org/zap"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
)

//投票

func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteDate)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}

		errDate := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errDate)
		return
	}
	//获取当前用户id
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}

	//投票业务逻辑
	err = logic.VoteForPost(userID, p)
	if err != nil {
		zap.L().Error("logic.VoteForPost failed", zap.Error(err))
		switch err {
		case redis.ErrVoteRepeated:
			ResponseError(c, ErrVoteRepeated)
		case redis.ErrVoteTimeExpire:
			ResponseError(c, ErrorVoteTimeExpire)
		default:
			ResponseError(c, CodeServerBusy)
		}
		return
	}
	ResponseSuccess(c, nil)
}
