package models

// 定义请求的参数结构体

// ParamSignUp 注册请求参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required,min=3,ne=admin"`
	Password   string `json:"password" binding:"required,min=6"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamPostList struct {
	Page int64 `json:"page" form:"page" example:"1"`  // 页码
	Size int64 `json:"size" form:"size" example:"10"` // 每页数据量
}

// 投票数据
type ParamVoteData struct {
	PostID string `json:"postid" binding:"required"`
}
