package controller

import (
	"gin_forum/logic"
	"gin_forum/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func GetPostListHandler(c *gin.Context) {

	page, size := getPageInfo(c)

	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// 返回响应
	ResponseSuccess(c, data)

}

func CreatePostHandler(c *gin.Context) {

	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误，直接返回响应

		zap.L().Error("post with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors 类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {

			ResponseError(c, CodeInvalidParam)
			return
		}

		ResponseErrorWithMsg(c, CodeInvalidParam, RemoveTopStruct(errs.Error()))
		return

	}
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	// fmt.Println("createpost")
	// fmt.Println(p.Title)
	// fmt.Println(p.Content)
	// fmt.Println(p.AuthorID)
	// fmt.Println("AuthorID 的数据类型是:", reflect.TypeOf(p.AuthorID))
	ResponseSuccess(c, gin.H{"msg": "create post success"})
}

func GetPostDetailHandler(c *gin.Context) {

	postid := c.Param("postid")

	post_id, err := strconv.ParseInt(postid, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}

	postdetail, err := logic.GetPostDetail(post_id)
	if err != nil {
		zap.L().Error("logic.GetPostDetail(p) failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)

		return
	}

	ResponseSuccess(c, postdetail)
}
