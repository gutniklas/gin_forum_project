package controller

import (
	"gin_forum/logic"
	"gin_forum/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type VoteData struct {
	PostID int64 `json:"postid,string"`
}

func PostVoteController(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {


		zap.L().Error("vote with invalid param", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Error()))
		return

	}
	userID, err := getCurrentUserID(c)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
	}
	PostScore, err := logic.VoteForPost(userID, p)
	if err != nil || PostScore < 0 {
		ResponseError(c, CodeServerBusy)
	}

	ResponseSuccess(c, gin.H{"Likenum": PostScore})

}
