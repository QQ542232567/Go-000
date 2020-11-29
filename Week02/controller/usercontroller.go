package controller

import (
	"errors"
	"strconv"
	"week002/dao"
	"week002/response"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	xerrors "github.com/pkg/errors"
)

type UserController struct {
}

func (c UserController) Show(ctx *gin.Context) {

	userid, _ := strconv.Atoi(ctx.Params.ByName("id"))

	userDao := dao.UserDao{}
	user, err := userDao.SelectById(userid)
	if err != nil {

		// 得到根因
		xerr := xerrors.Cause(err)

		// 依据根因，处理不同的逻辑？没有get到点。这错误wrap上来有什么好处呢
		if errors.Is(xerr, gorm.ErrRecordNotFound) {
			response.Fail(ctx, "xerr:"+err.Error(), nil)
		} else {
			response.Fail(ctx, "err:"+err.Error(), nil)

		}
		return
	}
	response.Success(ctx, gin.H{"user": user}, "success")
}
