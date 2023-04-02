package controller

import (
	"errors"
	"fmt"
	"gin_forum/dao/mysql"
	"gin_forum/logic"
	"gin_forum/models"

	"github.com/go-playground/validator/v10"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {

	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {


		zap.L().Error("SignUp with invalid param", zap.Error(err))
	
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Error()))
		return

	}


	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}

	ResponseSuccess(c, nil)
}
// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {

	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {


		zap.L().Error("SignUp with invalid param", zap.Error(err))
	
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Error()))
		return

	}


	user, err := logic.Login(p)

	if err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), 
		"user_name": user.Username,
		"token":     user.Token,
	})
}
